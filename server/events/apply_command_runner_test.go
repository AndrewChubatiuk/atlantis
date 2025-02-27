package events_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/core/locking"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/events/models/fixtures"
	"github.com/runatlantis/atlantis/server/logging"
	"github.com/runatlantis/atlantis/server/metrics"
)

func TestApplyCommandRunner_IsLocked(t *testing.T) {
	RegisterMockTestingT(t)

	cases := []struct {
		Description    string
		ApplyLocked    bool
		ApplyLockError error
		ExpComment     string
	}{
		{
			Description:    "When global apply lock is present IsDisabled returns true",
			ApplyLocked:    true,
			ApplyLockError: nil,
			ExpComment:     "**Error:** Running `atlantis apply` is disabled.",
		},
		{
			Description:    "When no global apply lock is present and DisableApply flag is false IsDisabled returns false",
			ApplyLocked:    false,
			ApplyLockError: nil,
			ExpComment:     "Ran Apply for 0 projects:\n\n\n\n",
		},
		{
			Description:    "If ApplyLockChecker returns an error IsDisabled return value of DisableApply flag",
			ApplyLockError: errors.New("error"),
			ApplyLocked:    false,
			ExpComment:     "Ran Apply for 0 projects:\n\n\n\n",
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			logger := logging.NewNoopCtxLogger(t)
			vcsClient := setup(t)

			scopeNull, _, _ := metrics.NewLoggingScope(logger, "atlantis")
			modelPull := models.PullRequest{BaseRepo: fixtures.GithubRepo, State: models.OpenPullState, Num: fixtures.Pull.Num}
			ctx := &command.Context{
				User:       fixtures.User,
				Log:        logger,
				Pull:       modelPull,
				HeadRepo:   fixtures.GithubRepo,
				Trigger:    command.CommentTrigger,
				Scope:      scopeNull,
				RequestCtx: context.TODO(),
			}

			When(applyLockChecker.CheckApplyLock()).ThenReturn(locking.ApplyCommandLock{Locked: c.ApplyLocked}, c.ApplyLockError)
			applyCommandRunner.Run(ctx, &command.Comment{Name: command.Apply})

			vcsClient.VerifyWasCalledOnce().CreateComment(fixtures.GithubRepo, modelPull.Num, c.ExpComment, "apply")
		})
	}
}
