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

func getCodeReview(prContents, gptModel string) {

	criticWindow.ProgressBar.StartProgressBar()

	review, err := critic.GetCodeReviewFromAPI(prContents, gptModel)

	if err != nil {
		ResetCenterStage()
		dialog.ShowError(fmt.Errorf("error getting review: %s", err), *criticWindow.Window)
		return
	}

	criticWindow.ReportPanel.Canvas.ParseMarkdown(review)
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

	err = critic.GetPullRequest(url, s, gptModel, prNumber, onGetPullRequestHandler)

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

	go getCodeReview(diff, gptModel)
}

func ResetCenterStage() {

	criticWindow.DiffPanel.Resize()
	criticWindow.ReportPanel.Resize()
	criticWindow.ReportPanel.Canvas.Scroll = container.ScrollBoth
	criticWindow.ProgressBar.StopProgressBar()

}
