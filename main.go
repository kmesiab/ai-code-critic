//go:generate ai-code-critic bundle bundled.go assets

package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
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
		onAPIKeySubmitButtonClickedHandler,
	)

	_, err := critic.GetConfig()

	if err != nil {
		dialog.ShowError(err, *criticWindow.Window)
	}

	(*criticWindow.Window).ShowAndRun()
}

func getCodeReview(prContents string) {

	resetPanel := func() {
		criticWindow.ProgressBar.Canvas.Stop()
		criticWindow.ProgressBar.Canvas.Hide()
		criticWindow.ReportPanel.Canvas.Scroll = container.ScrollBoth
		criticWindow.ReportPanel.Canvas.Resize(criticWindow.ReportPanel.Size)
	}

	review, err := critic.GetCodeReviewFromAPI(prContents)

	if err != nil {
		resetPanel()
		dialog.ShowError(fmt.Errorf("error getting review: %s", err), *criticWindow.Window)
		return
	}

	review = critic.ShortenLongLines(review, "\n")
	criticWindow.ReportPanel.Canvas.ParseMarkdown(review)
	resetPanel()
}

func onAPIKeySubmitButtonClickedHandler(ok bool) {

	if !ok {
		return
	}

	criticWindow.ProgressBar.Canvas.Start()
	criticWindow.ProgressBar.Canvas.Show()

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

	prContents = critic.ShortenLongLines(prContents, "\n\n")

	// Set the diff text
	criticWindow.DiffPanel.SetDiffText(prContents)

	// Set the report
	criticWindow.ReportPanel.Canvas.ParseMarkdown(critic.WaitingForReportMarkdown)
	criticWindow.ReportPanel.Canvas.Resize(criticWindow.ReportPanel.Size)

	// Send the pull request to the LLM
	getCodeReview(prContents)

}

func onFileOpenButtonClickedHandler() {
	criticWindow.PullRequestURLModal.Form.Show()
}

func onAnalyzeButtonClickedHandler() {

	windowSize := (*criticWindow.Window).Canvas().Size()
	halfSize := ShrinkByHalf(windowSize)

	criticWindow.ReportPanel.Canvas.Resize(halfSize)
	criticWindow.DiffPanel.Canvas.Resize(halfSize)

	(*criticWindow.Window).CenterOnScreen()
	(*criticWindow.Window).Resize(windowSize)

}

func ShrinkByHalf(size fyne.Size) fyne.Size {
	return fyne.NewSize(size.Width/2, size.Height/2)
}
