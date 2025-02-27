package runtime_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-getter"
	version "github.com/hashicorp/go-version"
	. "github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/core/runtime"
	"github.com/runatlantis/atlantis/server/core/runtime/mocks"
	tfMocks "github.com/runatlantis/atlantis/server/core/terraform/mocks"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/logging"
	. "github.com/runatlantis/atlantis/testing"
	"github.com/stretchr/testify/assert"
)

type NoopTFDownloader struct{}

func (m *NoopTFDownloader) GetFile(dst, src string, opts ...getter.ClientOption) error {
	return nil
}

func (m *NoopTFDownloader) GetAny(dst, src string, opts ...getter.ClientOption) error {
	return nil
}

func TestStepsRunner_Run(t *testing.T) {
	cases := []struct {
		description string
		steps       []valid.Step
		applyReqs   []string

		expSteps      []string
		expOut        string
		expFailure    string
		pullMergeable bool
	}{
		{
			description: "workflow with custom apply stage",
			steps: []valid.Step{
				{
					StepName:    "env",
					EnvVarName:  "key",
					EnvVarValue: "value",
				},
				{
					StepName: "run",
				},
				{
					StepName: "apply",
				},
				{
					StepName: "plan",
				},
				{
					StepName: "init",
				},
			},
			expSteps: []string{"env", "run", "apply", "plan", "init"},
			expOut:   "run\napply\nplan\ninit",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			RegisterMockTestingT(t)
			mockInit := mocks.NewMockRunner()
			mockPlan := mocks.NewMockRunner()
			mockShow := mocks.NewMockRunner()
			mockApply := mocks.NewMockRunner()
			mockRun := mocks.NewMockCustomRunner()
			mockEnv := mocks.NewMockEnvRunner()
			mockPolicyCheck := mocks.NewMockRunner()
			mockVersion := mocks.NewMockRunner()

			runner := runtime.NewStepsRunner(
				mockInit,
				mockPlan,
				mockShow,
				mockPolicyCheck,
				mockApply,
				mockVersion,
				mockRun,
				mockEnv,
			)
			repoDir, cleanup := TempDir(t)
			defer cleanup()

			ctx := context.Background()
			prjCtx := command.ProjectContext{
				Log:        logging.NewNoopCtxLogger(t),
				Steps:      c.steps,
				Workspace:  "default",
				RepoRelDir: ".",
			}
			expEnvs := map[string]string{
				"key": "value",
			}
			When(mockInit.Run(ctx, prjCtx, nil, repoDir, expEnvs)).ThenReturn("init", nil)
			When(mockPlan.Run(ctx, prjCtx, nil, repoDir, expEnvs)).ThenReturn("plan", nil)
			When(mockApply.Run(ctx, prjCtx, nil, repoDir, expEnvs)).ThenReturn("apply", nil)
			When(mockRun.Run(ctx, prjCtx, "", repoDir, expEnvs)).ThenReturn("run", nil)
			When(mockEnv.Run(ctx, prjCtx, "", "value", repoDir, make(map[string]string))).ThenReturn("value", nil)

			_, err := runner.Run(ctx, prjCtx, repoDir)
			assert.NoError(t, err)

			for _, step := range c.expSteps {
				switch step {
				case "init":
					mockInit.VerifyWasCalledOnce().Run(ctx, prjCtx, nil, repoDir, expEnvs)
				case "plan":
					mockPlan.VerifyWasCalledOnce().Run(ctx, prjCtx, nil, repoDir, expEnvs)
				case "apply":
					mockApply.VerifyWasCalledOnce().Run(ctx, prjCtx, nil, repoDir, expEnvs)
				case "run":
					mockRun.VerifyWasCalledOnce().Run(ctx, prjCtx, "", repoDir, expEnvs)
				case "env":
					mockEnv.VerifyWasCalledOnce().Run(ctx, prjCtx, "", "value", repoDir, expEnvs)
				}
			}
		})
	}
}

// Test run and env steps. We don't use mocks for this test since we're
// not running any Terraform.
func TestStepsRuinner_RunEnvSteps(t *testing.T) {
	RegisterMockTestingT(t)

	terraform := tfMocks.NewMockClient()
	tfVersion, err := version.NewVersion("0.12.0")
	Ok(t, err)
	mockInit := mocks.NewMockRunner()
	mockPlan := mocks.NewMockRunner()
	mockShow := mocks.NewMockRunner()
	mockApply := mocks.NewMockRunner()
	mockPolicyCheck := mocks.NewMockRunner()
	mockVersion := mocks.NewMockRunner()

	run := &runtime.RunStepRunner{
		TerraformExecutor: terraform,
		DefaultTFVersion:  tfVersion,
	}

	runner := runtime.NewStepsRunner(
		mockInit,
		mockPlan,
		mockShow,
		mockPolicyCheck,
		mockApply,
		mockVersion,
		run,
		&runtime.EnvStepRunner{
			RunStepRunner: run,
		},
	)

	repoDir, cleanup := TempDir(t)
	defer cleanup()

	ctx := context.Background()
	prjCtx := command.ProjectContext{
		Log:        logging.NewNoopCtxLogger(t),
		RequestCtx: context.TODO(),
		Steps: []valid.Step{
			{
				StepName:   "run",
				RunCommand: "echo var=$var",
			},
			{
				StepName:    "env",
				EnvVarName:  "var",
				EnvVarValue: "value",
			},
			{
				StepName:   "run",
				RunCommand: "echo var=$var",
			},
			{
				StepName:   "env",
				EnvVarName: "dynamic_var",
				RunCommand: "echo dynamic_value",
			},
			{
				StepName:   "run",
				RunCommand: "echo dynamic_var=$dynamic_var",
			},
			// Test overriding the variable
			{
				StepName:    "env",
				EnvVarName:  "dynamic_var",
				EnvVarValue: "overridden",
			},
			{
				StepName:   "run",
				RunCommand: "echo dynamic_var=$dynamic_var",
			},
		},
		Workspace:  "default",
		RepoRelDir: ".",
	}
	res, err := runner.Run(ctx, prjCtx, repoDir)
	Ok(t, err)

	Equals(t, "var=\n\nvar=value\n\ndynamic_var=dynamic_value\n\ndynamic_var=overridden\n", res)
}
