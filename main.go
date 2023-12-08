//go:generate ai-code-critic bundle bundled.go assets

package main

import (
	"fmt"
	"strconv"

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

func getCodeReview(prContents string) {

	ResetCenterStage()

	review, err := critic.GetCodeReviewFromAPI(prContents)

	if err != nil {
		ResetCenterStage()
		dialog.ShowError(fmt.Errorf("error getting review: %s", err), *criticWindow.Window)
		return
	}

	review = critic.ShortenLongLines(review, "\n")
	criticWindow.ReportPanel.Canvas.ParseMarkdown(review)
	ResetCenterStage()
}

func onPullRequestModalClickedHandler(ok bool) {

	if !ok {
		return
	}

	input := criticWindow.PullRequestURLModal.TextEntry.Text

	if input == "" {
		return
	}

	criticWindow.ProgressBar.Canvas.Start()
	criticWindow.ProgressBar.Canvas.Show()

	url, s, s2, err := critic.ParseGithubPullRequestURL(input)

	if err != nil {
		dialog.ShowError(fmt.Errorf("error parsing URL: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}

	prNumber, err := strconv.Atoi(s2)

	if err != nil {
		dialog.ShowError(fmt.Errorf("invalid PR number: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}

	err = critic.GetPullRequest(url, s, prNumber, onGetPullRequestHandler)

	if err != nil {
		dialog.ShowError(fmt.Errorf("error getting PR: %s", err), *criticWindow.Window)
		ResetCenterStage()

		return
	}
}

func onGetPullRequestHandler(prContents string) {

	// prContents = critic.ShortenLongLines(prContents, "\n\n")

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

	ResetCenterStage()
	(*criticWindow.Window).CenterOnScreen()
}

func ResetCenterStage() {

	criticWindow.DiffPanel.Canvas.Resize(criticWindow.DiffPanel.Canvas.Size())

	criticWindow.ProgressBar.Canvas.Stop()
	criticWindow.ProgressBar.Canvas.Hide()

	criticWindow.ReportPanel.Canvas.Resize(criticWindow.ReportPanel.Canvas.Size())
	criticWindow.ReportPanel.Canvas.Scroll = container.ScrollBoth

}
