package internal

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

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

func GetPullRequest(owner string, repo string, prNumber int, callback OnGetPullRequestEvent) error {
	timeout, err := GetContextTimeout()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client := github.NewClient(nil)
	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repo, prNumber)
	if err != nil {
		return err
	}

	pullRequest.GetDiffURL()

	ch := make(chan string)
	go getDiffContents(ch, pullRequest.GetDiffURL())

	contents := <-ch

	callback(contents)

	return nil
}

func getDiffContents(c chan<- string, diffURL string) {
	diffContents, err := http.Get(diffURL)
	if err != nil {
		c <- err.Error()
		return
	}

	bodyBytes, err := io.ReadAll(diffContents.Body)
	if err != nil {
		c <- err.Error()
		return
	}

	c <- string(bodyBytes)
}

func GetContextTimeout() (time.Duration, error) {
	cfg, err := GetConfig()
	if err != nil {
		return 0, err
	}

	return cfg.ContextTimeout, nil
}
