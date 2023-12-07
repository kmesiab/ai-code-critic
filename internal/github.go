package internal

import (
	"context"

	"github.com/google/go-github/v57/github"
)

func GetPullRequest(owner string, repo string, prNumber int) (string, error) {

	ctx := context.Background()
	client := github.NewClient(nil)
	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repo, prNumber)

	if err != nil {
		return "", err
	}

	return *pullRequest.Body, nil
}
