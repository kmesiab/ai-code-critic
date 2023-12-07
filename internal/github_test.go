package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGithub_ParseGithubPullRequestURL(t *testing.T) {
	owner, repo, prNumber, err := ParseGithubPullRequestURL(
		"https://github.com/google/go-github/pull/1234",
	)

	require.NoError(t, err)

	if owner != "google" || repo != "go-github" || prNumber != "1234" {
		t.Error("failed to parse pull request URL")
	}
}

func TestGithub_ParseGithubPullRequestInvalidURL(t *testing.T) {
	owner, repo, prNumber, err := ParseGithubPullRequestURL("foo")

	require.Error(t, err)
	require.Emptyf(t, owner, "owner should be empty")
	require.Emptyf(t, repo, "repo should be empty")
	require.Emptyf(t, prNumber, "prNumber should be empty")
}
