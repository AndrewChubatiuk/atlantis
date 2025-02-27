package raw

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/runatlantis/atlantis/server/core/config/valid"
)

type Workflows map[string]Workflow

func (w Workflows) ToValid(defaultCfg valid.GlobalCfg) map[string]valid.Workflow {
	validWorkflows := make(map[string]valid.Workflow)
	for k, v := range w {
		validWorkflows[k] = v.ToValid(k)
	}

	// Merge in defaults without overriding.
	for k, v := range defaultCfg.Workflows {
		if _, ok := validWorkflows[k]; !ok {
			validWorkflows[k] = v
		}
	}

	return validWorkflows
}

type Workflow struct {
	Apply       *Stage `yaml:"apply,omitempty" json:"apply,omitempty"`
	Plan        *Stage `yaml:"plan,omitempty" json:"plan,omitempty"`
	PolicyCheck *Stage `yaml:"policy_check,omitempty" json:"policy_check,omitempty"`
}

func (w Workflow) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Apply),
		validation.Field(&w.Plan),
		validation.Field(&w.PolicyCheck),
	)
}

func (w Workflow) toValidStage(stage *Stage, defaultStage valid.Stage) valid.Stage {
	if stage == nil || stage.Steps == nil {
		return defaultStage
	}

	return stage.ToValid()
}

func (w Workflow) ToValid(name string) valid.Workflow {
	v := valid.Workflow{
		Name: name,
	}

	v.Apply = w.toValidStage(w.Apply, valid.DefaultApplyStage)
	v.Plan = w.toValidStage(w.Plan, valid.DefaultPlanStage)
	v.PolicyCheck = w.toValidStage(w.PolicyCheck, valid.DefaultPolicyCheckStage)

	return v
}
