//go:generate ai-code-critic bundle bundled.go assets

package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	critic "github.com/kmesiab/ai-code-critic/internal"
	"github.com/kmesiab/ai-code-critic/ui"
)

var criticWindow *ui.CriticWindow

func main() {
	application := app.New()

	criticWindow = ui.Initialize(application,
		onFileOpenButtonClickedHandler,
		onAnalyzeButtonClickedHandler,
		onAPIKeySubmitButtonClickedHandler,
	)

	(*criticWindow.Window).ShowAndRun()

}

func getCodeReview(prContents string) {
	review, err := critic.GetCodeReviewFromAPI(prContents)

	if err != nil {
		critic.Logf("Error getting review: %s", err)
	}

	criticWindow.ReportPanel.Canvas.ParseMarkdown(review)

	// Resize the window to half the size of the parent container
	canvasSize := (*criticWindow.Canvas).Size()
	newSize := fyne.NewSize(
		canvasSize.Width/2, criticWindow.ReportPanel.Canvas.Size().Height,
	)

	criticWindow.ProgressBar.Canvas.Stop()
	criticWindow.ReportPanel.Canvas.Resize(newSize)
}

func onAPIKeySubmitButtonClickedHandler(ok bool) {

	criticWindow.ProgressBar.Canvas.Start()

	if !ok {
		return
	}

	input := criticWindow.PullRequestURLModal.TextEntry.Text

	if input == "" {
		return
	}

	url, s, s2, err := critic.ParseGithubPullRequestURL(input)

	if err != nil {
		critic.Logf("Error parsing URL: %s", err)
	}

	prNumber, err := strconv.Atoi(s2)

	if err != nil {
		critic.Logf("Invalid PR number: %s", s2)
	}

	err = critic.GetPullRequest(url, s, prNumber, onGetPullRequestHandler)

	if err != nil {
		critic.Logf("Error getting PR: %s", err)
	}
}

func onGetPullRequestHandler(prContents string) {

	// Set the diff text
	criticWindow.DiffPanel.SetText(prContents)

	// Set the report
	criticWindow.ReportPanel.Canvas.ParseMarkdown(critic.WaitingForReportMarkdown)

	// Send the pull request to the LLM
	getCodeReview(prContents)

}

func onFileOpenButtonClickedHandler() {
	criticWindow.PullRequestURLModal.Form.Show()
}

func onAnalyzeButtonClickedHandler() {

	criticWindow.ReportPanel.Canvas.Resize(ShrinkByHalf(criticWindow.ReportPanel.Size))
	criticWindow.DiffPanel.Canvas.Resize(ShrinkByHalf(criticWindow.DiffPanel.Size))
	(*criticWindow.Window).CenterOnScreen()
	(*criticWindow.Window).Resize(criticWindow.Size)

}

func ShrinkByHalf(size fyne.Size) fyne.Size {
	return fyne.NewSize(size.Width/2, size.Height/2)
}
