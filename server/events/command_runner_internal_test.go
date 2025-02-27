package events

import (
	"context"
	"fmt"
	"testing"

	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/events/models"
	. "github.com/runatlantis/atlantis/testing"
)

func TestApplyUpdateCommitStatus(t *testing.T) {
	cases := map[string]struct {
		cmd           command.Name
		pullStatus    models.PullStatus
		expStatus     models.VCSStatus
		expNumSuccess int
		expNumTotal   int
	}{
		"apply, one pending": {
			cmd: command.Apply,
			pullStatus: models.PullStatus{
				Projects: []models.ProjectStatus{
					{
						Status: models.PlannedPlanStatus,
					},
					{
						Status: models.AppliedPlanStatus,
					},
				},
			},
			expStatus:     models.PendingVCSStatus,
			expNumSuccess: 1,
			expNumTotal:   2,
		},
		"apply, all successful": {
			cmd: command.Apply,
			pullStatus: models.PullStatus{
				Projects: []models.ProjectStatus{
					{
						Status: models.AppliedPlanStatus,
					},
					{
						Status: models.AppliedPlanStatus,
					},
				},
			},
			expStatus:     models.SuccessVCSStatus,
			expNumSuccess: 2,
			expNumTotal:   2,
		},
		"apply, one errored, one pending": {
			cmd: command.Apply,
			pullStatus: models.PullStatus{
				Projects: []models.ProjectStatus{
					{
						Status: models.AppliedPlanStatus,
					},
					{
						Status: models.ErroredApplyStatus,
					},
					{
						Status: models.PlannedPlanStatus,
					},
				},
			},
			expStatus:     models.FailedVCSStatus,
			expNumSuccess: 1,
			expNumTotal:   3,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			csu := &MockCSU{}
			cr := &ApplyCommandRunner{
				vcsStatusUpdater: csu,
			}
			cr.updateVcsStatus(&command.Context{}, c.pullStatus, "")
			Equals(t, models.Repo{}, csu.CalledRepo)
			Equals(t, models.PullRequest{}, csu.CalledPull)
			Equals(t, c.expStatus, csu.CalledStatus)
			Equals(t, c.cmd.String(), csu.CalledCommand)
			Equals(t, c.expNumSuccess, csu.CalledNumSuccess)
			Equals(t, c.expNumTotal, csu.CalledNumTotal)
		})
	}
}

func TestPlanUpdateCommitStatus(t *testing.T) {
	cases := map[string]struct {
		cmd           command.Name
		pullStatus    models.PullStatus
		expStatus     models.VCSStatus
		expNumSuccess int
		expNumTotal   int
	}{
		"single plan success": {
			cmd: command.Plan,
			pullStatus: models.PullStatus{
				Projects: []models.ProjectStatus{
					{
						Status: models.PlannedPlanStatus,
					},
				},
			},
			expStatus:     models.SuccessVCSStatus,
			expNumSuccess: 1,
			expNumTotal:   1,
		},
		"one plan error, other errors": {
			cmd: command.Plan,
			pullStatus: models.PullStatus{
				Projects: []models.ProjectStatus{
					{
						Status: models.ErroredPlanStatus,
					},
					{
						Status: models.PlannedPlanStatus,
					},
					{
						Status: models.AppliedPlanStatus,
					},
					{
						Status: models.ErroredApplyStatus,
					},
				},
			},
			expStatus:     models.FailedVCSStatus,
			expNumSuccess: 3,
			expNumTotal:   4,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			csu := &MockCSU{}
			cr := &PlanCommandRunner{
				vcsStatusUpdater: csu,
			}
			cr.updateVcsStatus(&command.Context{}, c.pullStatus, "")
			Equals(t, models.Repo{}, csu.CalledRepo)
			Equals(t, models.PullRequest{}, csu.CalledPull)
			Equals(t, c.expStatus, csu.CalledStatus)
			Equals(t, c.cmd.String(), csu.CalledCommand)
			Equals(t, c.expNumSuccess, csu.CalledNumSuccess)
			Equals(t, c.expNumTotal, csu.CalledNumTotal)
		})
	}
}

type MockCSU struct {
	CalledRepo       models.Repo
	CalledPull       models.PullRequest
	CalledStatus     models.VCSStatus
	CalledCommand    string
	CalledNumSuccess int
	CalledNumTotal   int
	CalledStatusID   string
}

func (m *MockCSU) UpdateCombinedCount(ctx context.Context, repo models.Repo, pull models.PullRequest, status models.VCSStatus, command fmt.Stringer, numSuccess int, numTotal int, statusID string) (string, error) {
	m.CalledRepo = repo
	m.CalledPull = pull
	m.CalledStatus = status
	m.CalledCommand = command.String()
	m.CalledNumSuccess = numSuccess
	m.CalledNumTotal = numTotal
	m.CalledStatusID = statusID
	return "", nil
}
func (m *MockCSU) UpdateCombined(ctx context.Context, repo models.Repo, pull models.PullRequest, status models.VCSStatus, command fmt.Stringer, statusID string, output string) (string, error) {
	return "", nil
}
func (m *MockCSU) UpdateProject(ctx context.Context, projectCtx command.ProjectContext, cmdName fmt.Stringer, status models.VCSStatus, url string, statusID string) (string, error) {
	return "", nil
}
