package command

import (
	"github.com/runatlantis/atlantis/server/events/models"
)

// ProjectResult is the result of executing a plan/policy_check/apply for a specific project.
type ProjectResult struct {
	Command            Name
	RepoRelDir         string
	Workspace          string
	Error              error
	Failure            string
	PlanSuccess        *models.PlanSuccess
	PolicyCheckSuccess *models.PolicyCheckSuccess
	ApplySuccess       string
	VersionSuccess     string
	ProjectName        string
	StatusID           string
	JobID              string
}

// VcsStatus returns the vcs commit status of this project result.
func (p ProjectResult) VcsStatus() models.VCSStatus {
	if p.Error != nil {
		return models.FailedVCSStatus
	}
	if p.Failure != "" {
		return models.FailedVCSStatus
	}
	return models.SuccessVCSStatus
}

// PlanStatus returns the plan status.
func (p ProjectResult) PlanStatus() models.ProjectPlanStatus {
	switch p.Command {
	case Plan:
		if p.Error != nil {
			return models.ErroredPlanStatus
		} else if p.Failure != "" {
			return models.ErroredPlanStatus
		}
		return models.PlannedPlanStatus
	case PolicyCheck:
		if p.Error != nil {
			return models.ErroredPolicyCheckStatus
		} else if p.Failure != "" {
			return models.ErroredPolicyCheckStatus
		}
		return models.PassedPolicyCheckStatus
	case Apply:
		if p.Error != nil {
			return models.ErroredApplyStatus
		} else if p.Failure != "" {
			return models.ErroredApplyStatus
		}
		return models.AppliedPlanStatus
	}

	panic("PlanStatus() missing a combination")
}

// IsSuccessful returns true if this project result had no errors.
func (p ProjectResult) IsSuccessful() bool {
	return p.PlanSuccess != nil || p.PolicyCheckSuccess != nil || p.ApplySuccess != ""
}
