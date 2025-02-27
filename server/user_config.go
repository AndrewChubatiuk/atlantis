package server

import (
	"github.com/runatlantis/atlantis/server/logging"
)

type Mode int

const (
	Default Mode = iota
	Gateway
	Worker
	TemporalWorker
)

// UserConfig holds config values passed in by the user.
// The mapstructure tags correspond to flags in cmd/server.go and are used when
// the config is parsed from a YAML file.
type UserConfig struct {
	AtlantisURL                string `mapstructure:"atlantis-url"`
	AutoplanFileList           string `mapstructure:"autoplan-file-list"`
	AzureDevopsToken           string `mapstructure:"azuredevops-token"`
	AzureDevopsUser            string `mapstructure:"azuredevops-user"`
	AzureDevopsWebhookPassword string `mapstructure:"azuredevops-webhook-password"`
	AzureDevopsWebhookUser     string `mapstructure:"azuredevops-webhook-user"`
	BitbucketBaseURL           string `mapstructure:"bitbucket-base-url"`
	BitbucketToken             string `mapstructure:"bitbucket-token"`
	BitbucketUser              string `mapstructure:"bitbucket-user"`
	BitbucketWebhookSecret     string `mapstructure:"bitbucket-webhook-secret"`
	CheckoutStrategy           string `mapstructure:"checkout-strategy"`
	DataDir                    string `mapstructure:"data-dir"`
	DisableApplyAll            bool   `mapstructure:"disable-apply-all"`
	DisableApply               bool   `mapstructure:"disable-apply"`
	DisableAutoplan            bool   `mapstructure:"disable-autoplan"`
	DisableMarkdownFolding     bool   `mapstructure:"disable-markdown-folding"`
	EnablePolicyChecks         bool   `mapstructure:"enable-policy-checks"`
	EnableRegExpCmd            bool   `mapstructure:"enable-regexp-cmd"`
	EnableDiffMarkdownFormat   bool   `mapstructure:"enable-diff-markdown-format"`
	FFOwner                    string `mapstructure:"ff-owner"`
	FFRepo                     string `mapstructure:"ff-repo"`
	FFBranch                   string `mapstructure:"ff-branch"`
	FFPath                     string `mapstructure:"ff-path"`
	GithubHostname             string `mapstructure:"gh-hostname"`
	GithubToken                string `mapstructure:"gh-token"`
	GithubUser                 string `mapstructure:"gh-user"`
	GithubWebhookSecret        string `mapstructure:"gh-webhook-secret"`
	GithubOrg                  string `mapstructure:"gh-org"`
	GithubAppID                int64  `mapstructure:"gh-app-id"`
	GithubAppKey               string `mapstructure:"gh-app-key"`
	GithubAppKeyFile           string `mapstructure:"gh-app-key-file"`
	GithubAppSlug              string `mapstructure:"gh-app-slug"`
	GitlabHostname             string `mapstructure:"gitlab-hostname"`
	GitlabToken                string `mapstructure:"gitlab-token"`
	GitlabUser                 string `mapstructure:"gitlab-user"`
	GitlabWebhookSecret        string `mapstructure:"gitlab-webhook-secret"`
	HidePrevPlanComments       bool   `mapstructure:"hide-prev-plan-comments"`
	LogLevel                   string `mapstructure:"log-level"`
	ParallelPoolSize           int    `mapstructure:"parallel-pool-size"`
	MaxProjectsPerPR           int    `mapstructure:"max-projects-per-pr"`
	StatsNamespace             string `mapstructure:"stats-namespace"`
	PlanDrafts                 bool   `mapstructure:"allow-draft-prs"`
	Port                       int    `mapstructure:"port"`
	RepoConfig                 string `mapstructure:"repo-config"`
	RepoConfigJSON             string `mapstructure:"repo-config-json"`
	RepoAllowlist              string `mapstructure:"repo-allowlist"`
	// RepoWhitelist is deprecated in favour of RepoAllowlist.
	RepoWhitelist string `mapstructure:"repo-whitelist"`

	// RequireUnDiverged is whether to require pull requests to rebase default branch before
	// allowing terraform apply's to run.
	RequireUnDiverged bool `mapstructure:"require-undiverged"`
	// RequireSQUnlocked is whether to require pull requests to be unlocked before running
	// terraform apply.
	RequireSQUnlocked        bool            `mapstructure:"require-unlocked"`
	SlackToken               string          `mapstructure:"slack-token"`
	SSLCertFile              string          `mapstructure:"ssl-cert-file"`
	SSLKeyFile               string          `mapstructure:"ssl-key-file"`
	TFDownloadURL            string          `mapstructure:"tf-download-url"`
	VCSStatusName            string          `mapstructure:"vcs-status-name"`
	DefaultTFVersion         string          `mapstructure:"default-tf-version"`
	Webhooks                 []WebhookConfig `mapstructure:"webhooks"`
	WriteGitCreds            bool            `mapstructure:"write-git-creds"`
	LyftAuditJobsSnsTopicArn string          `mapstructure:"lyft-audit-jobs-sns-topic-arn"`
	LyftGatewaySnsTopicArn   string          `mapstructure:"lyft-gateway-sns-topic-arn"`
	LyftMode                 string          `mapstructure:"lyft-mode"`
	LyftWorkerQueueURL       string          `mapstructure:"lyft-worker-queue-url"`

	// Supports adding a default URL to the checkrun UI when details URL is not set
	DefaultCheckrunDetailsURL string `mapstructure:"default-checkrun-details-url"`
}

// ToLogLevel returns the LogLevel object corresponding to the user-passed
// log level.
func (u UserConfig) ToLogLevel() logging.LogLevel {
	switch u.LogLevel {
	case "debug":
		return logging.Debug
	case "info":
		return logging.Info
	case "warn":
		return logging.Warn
	case "error":
		return logging.Error
	}
	return logging.Info
}

// ToLyftMode returns mode type to run atlantis on.
func (u UserConfig) ToLyftMode() Mode {
	switch u.LyftMode {
	case "default":
		return Default
	case "gateway":
		return Gateway
	case "worker":
		return Worker
	case "temporalworker":
		return TemporalWorker
	}
	return Default
}
