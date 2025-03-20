package vcs

// GithubConfig allows for custom github-specific functionality and behavior
type GithubConfig struct {
	AllowMergeableBypassApply bool
	// MaxCommentLength is the maximum length of GitHub comments that Atlantis will create.
	// Comments longer than this will be split into multiple comments.
	// If not specified, defaults to 65536.
	MaxCommentLength int `mapstructure:"max_comment_length"`
}
