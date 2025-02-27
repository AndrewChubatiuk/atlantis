package converter

import (
	"fmt"
	"regexp"

	"github.com/google/go-github/v45/github"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/neptune/gateway/event"
	"github.com/runatlantis/atlantis/server/vcs"
)

var refRegex = regexp.MustCompile("refs/(?P<type>(?:heads)|(?:tags))/(?P<name>.+)")

// PushEvent converts a github pull request to our internal model
type PushEvent struct {
	RepoConverter RepoConverter
}

func (p PushEvent) Convert(e *github.PushEvent) (event.Push, error) {
	repo, err := p.RepoConverter.Convert(e.Repo)

	if err != nil {
		return event.Push{}, errors.Wrap(err, "converting repo")
	}

	matches := refRegex.FindStringSubmatch(e.GetRef())

	if len(matches) != 3 {
		return event.Push{}, fmt.Errorf("unable to determine ref")
	}

	t := matches[refRegex.SubexpIndex("type")]
	name := matches[refRegex.SubexpIndex("name")]

	refType, err := vcs.FromGithubRefType(t)

	if err != nil {
		return event.Push{}, errors.Wrap(err, "getting ref type")
	}

	installationToken := githubapp.GetInstallationIDFromEvent(e)

	action := event.UpdatedAction

	if e.GetCreated() {
		action = event.CreatedAction
	}

	if e.GetDeleted() {
		action = event.DeletedAction
	}

	return event.Push{
		Repo:   repo,
		Sha:    e.GetHeadCommit().GetID(),
		Action: action,
		Sender: models.User{
			Username: e.GetSender().GetLogin(),
		},
		Ref: vcs.Ref{
			Type: refType,
			Name: name,
		},
		InstallationToken: installationToken,
	}, nil
}
