package internal

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-github/v57/github"
)

func ParseGithubPullRequestURL(pullRequestURL string) (string, string, string, error) {

	parts := strings.Split(pullRequestURL, "/")

	if len(parts) != 7 {
		return "", "", "", errors.New("invalid pull request URL")
	}

	owner := parts[3]
	repo := parts[4]
	prNumber := parts[6]

	return owner, repo, prNumber, nil
}

func GetPullRequest(owner string, repo string, prNumber int) (string, error) {

	ctx := context.Background()
	client := github.NewClient(nil)
	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repo, prNumber)

	if err != nil {
		return "", err
	}

	pullRequest.GetDiffURL()
	diffContents, err := http.Get(pullRequest.GetDiffURL())

	if err != nil {
		return "", err
	}

	bodyBytes, err := io.ReadAll(diffContents.Body)

	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
