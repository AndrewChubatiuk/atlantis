package raw_test

import (
	"testing"

	"github.com/runatlantis/atlantis/server/core/config/raw"
	"github.com/runatlantis/atlantis/server/core/config/valid"
	. "github.com/runatlantis/atlantis/testing"
	yaml "gopkg.in/yaml.v2"
)

func TestPreWorkflowHook_YAMLMarshalling(t *testing.T) {
	cases := []struct {
		description string
		input       string
		exp         raw.PreWorkflowHook
		expErr      string
	}{
		// Run-step style
		{
			description: "run step",
			input: `
run: my command`,
			exp: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"run": "my command",
				},
			},
		},
		{
			description: "run step multiple top-level keys",
			input: `
run: my command
key: value`,
			exp: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"run": "my command",
					"key": "value",
				},
			},
		},

		// Errors
		{
			description: "extra args style no slice strings",
			input: `
key:
  value:
    another: map`,
			expErr: "yaml: unmarshal errors:\n  line 3: cannot unmarshal !!map into string",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			var got raw.PreWorkflowHook
			err := yaml.UnmarshalStrict([]byte(c.input), &got)
			if c.expErr != "" {
				ErrEquals(t, c.expErr, err)
				return
			}
			Ok(t, err)
			Equals(t, c.exp, got)

			_, err = yaml.Marshal(got)
			Ok(t, err)

			var got2 raw.PreWorkflowHook
			err = yaml.UnmarshalStrict([]byte(c.input), &got2)
			Ok(t, err)
			Equals(t, got2, got)
		})
	}
}

func TestGlobalConfigStep_Validate(t *testing.T) {
	cases := []struct {
		description string
		input       raw.PreWorkflowHook
		expErr      string
	}{
		{
			description: "run step",
			input: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"run": "my command",
				},
			},
			expErr: "",
		},
		{
			description: "invalid key in string val",
			input: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"invalid": "",
				},
			},
			expErr: "\"invalid\" is not a valid step type",
		},
		{
			// For atlantis.yaml v2, this wouldn't parse, but now there should
			// be no error.
			description: "unparseable shell command",
			input: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"run": "my 'c",
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			err := c.input.Validate()
			if c.expErr == "" {
				Ok(t, err)
				return
			}
			ErrEquals(t, c.expErr, err)
		})
	}
}

func TestPreWorkflowHook_ToValid(t *testing.T) {
	cases := []struct {
		description string
		input       raw.PreWorkflowHook
		exp         *valid.PreWorkflowHook
	}{
		{
			description: "run step",
			input: raw.PreWorkflowHook{
				StringVal: map[string]string{
					"run": "my 'run command'",
				},
			},
			exp: &valid.PreWorkflowHook{
				StepName:   "run",
				RunCommand: "my 'run command'",
			},
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			Equals(t, c.exp, c.input.ToValid())
		})
	}
}
