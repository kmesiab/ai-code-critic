package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	critic "github.com/kmesiab/ai-code-critic/internal"
	"github.com/kmesiab/ai-code-critic/ui"
)

var criticWindow *ui.CriticWindow

func main() {
	application := app.New()
	// application.Settings().SetTheme(theme.DarkTheme())

	criticWindow = ui.Initialize(application,
		onFileOpenButtonClickedHandler,
		onAnalyzeButtonClickedHandler,
		onPullRequestModalClickedHandler,
	)

	_, err := critic.GetConfig()
	if err != nil {
		dialog.ShowError(err, *criticWindow.Window)
	}

	(*criticWindow.Window).ShowAndRun()
}

func getCodeReview(prContents, gptModel string) {
	criticWindow.ProgressBar.StartProgressBar()

	config, err := critic.GetConfig()

	if err != nil {
		dialog.ShowError(fmt.Errorf("error getting config: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}

	ignoreList := strings.Split(config.IgnoreFiles, ",")

	// Step 1: Split and parse PR contents into individual GitDiff objects.
	gitDiffs := critic.ParseGitDiff(prContents, ignoreList)

	// Channel to collect responses from goroutines.
	reviewsChan := make(chan string, len(gitDiffs))
	var wg sync.WaitGroup

	for _, gitDiff := range gitDiffs {
		wg.Add(1)
		go func(gd *critic.GitDiff) {
			defer wg.Done()

			// Step 2: Send the diff to GetCodeReviewFromAPI in a goroutine.
			review, err := critic.GetCodeReviewFromAPI(gd.DiffContents, gptModel)
			if err != nil {
				reviewsChan <- fmt.Sprintf("Error getting review for %s: %s\n", gd.FilePathNew, err)
				return
			}

			// Prepend the diff header to the review.
			header := fmt.Sprintf("Review for %s:\n", gd.FilePathNew)
			reviewsChan <- header + review
		}(gitDiff)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(reviewsChan)

	// Step 4: Assemble all responses into a single string.
	var fullReview strings.Builder
	for review := range reviewsChan {
		fullReview.WriteString(review + "\n\n")
	}

	criticWindow.ReportPanel.Canvas.ParseMarkdown(fullReview.String())
	ResetCenterStage()
}

func onPullRequestModalClickedHandler(ok bool) {
	if !ok {
		return
	}

	input := criticWindow.PullRequestURLModal.TextEntry.Text
	gptModel := criticWindow.PullRequestURLModal.GPTModel.Text

	if input == "" || gptModel == "" {
		return
	}

	criticWindow.ProgressBar.Canvas.Start()
	criticWindow.ProgressBar.Canvas.Show()

	owner, repo, prNumber, err := critic.ParseGithubPullRequestURL(input)
	if err != nil {
		dialog.ShowError(fmt.Errorf("error parsing URL: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}

	prNumberInt, err := strconv.Atoi(prNumber)
	if err != nil {
		dialog.ShowError(fmt.Errorf("invalid PR number: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}

	err = critic.GetPullRequest(owner, repo, gptModel, prNumberInt, onGetPullRequestHandler)

	if err != nil {
		dialog.ShowError(fmt.Errorf("error getting PR: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}
}

func onGetPullRequestHandler(prContents, gptModel string) {

	// Set the diff text
	criticWindow.DiffPanel.SetDiffText(prContents)

	// Set the report
	criticWindow.ReportPanel.Canvas.ParseMarkdown(critic.WaitingForReportMarkdown)

	// Send the pull request to the LLM
	go getCodeReview(prContents, gptModel)
}

func onFileOpenButtonClickedHandler() {
	criticWindow.PullRequestURLModal.TextEntry.Text = ""
	criticWindow.PullRequestURLModal.GPTModel.Text = "Select one or type"
	criticWindow.PullRequestURLModal.Form.Show()
}

func onAnalyzeButtonClickedHandler() {
	// Prevent re-entry if already processing
	if criticWindow.IsAnalyzeButtonActive {
		dialog.ShowInformation("Analyze in Progress", "Analysis is already in progress. Please wait until the current process is complete.", *criticWindow.Window)

		return
	}

	diff := criticWindow.DiffPanel.TextGrid.Text()
	gptModel := criticWindow.PullRequestURLModal.GPTModel.Text

	if diff == "" || criticWindow.DiffPanel.IsDefaultText() {
		dialog.ShowError(fmt.Errorf("the diff is empty"), *criticWindow.Window)

		return
	}
	if gptModel == "" {
		dialog.ShowError(fmt.Errorf("no gpt model specified"), *criticWindow.Window)

		return
	}

	criticWindow.IsAnalyzeButtonActive = true

	go func() {
		getCodeReview(diff, gptModel)
		criticWindow.IsAnalyzeButtonActive = false
	}()
}

func ResetCenterStage() {
	criticWindow.DiffPanel.Resize()
	criticWindow.ReportPanel.Resize()
	criticWindow.ReportPanel.Canvas.Scroll = container.ScrollBoth
	criticWindow.ProgressBar.StopProgressBar()
}
