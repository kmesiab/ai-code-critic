package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	critic "github.com/kmesiab/ai-code-critic/internal"
	"github.com/kmesiab/ai-code-critic/ui/components"
)

func Initialize(
	app fyne.App,
	fileOpenButtonClickedHandler critic.FileOpenClickedEventHandler,
	analyzeButtonClickedHandler critic.AnalyzeButtonClickedEventHandler,
	submitButtonClickedEventHandler critic.SubmitButtonClickedEventHandler,
) *CriticWindow {

	canvasSize := fyne.NewSize(critic.MainCanvasWidth, critic.MainCanvasHeight)

	// Holds the pull request diff
	diffPanel := components.NewDiffPanel(canvasSize, "")

	// Holds the code review in Markdown format
	reportPanel := components.NewReportPanel(canvasSize, critic.IntroMarkdown)

	// Holds the progress bar used when fetching the diff and review
	progressBar := components.NewProgressBar(canvasSize)

	// The toolbar exposes two buttons, one for the modal to enter a pull
	// request url, and another to analyze the contents of the diffPanel
	toolbar := components.NewToolBar(
		fileOpenButtonClickedHandler,
		analyzeButtonClickedHandler,
	)

	centerStage := container.NewHSplit(reportPanel.Canvas, diffPanel.Canvas)
	mainStage := container.NewBorder(toolbar, progressBar.Canvas, nil, nil, centerStage)

	// Main program window
	window := app.NewWindow(critic.ApplicationName)
	window.SetContent(mainStage)
	window.Resize(canvasSize)

	// Create the pull requests url modal
	PullRequestURLModal := components.NewPullRequestURLModal(
		canvasSize, critic.PullRequestURLModalDefaultText,
		&window, submitButtonClickedEventHandler,
	)

	criticWindow := &CriticWindow{
		App:                 &app,
		Size:                canvasSize,
		ReportPanel:         reportPanel,
		DiffPanel:           diffPanel,
		ToolBar:             &toolbar,
		Canvas:              mainStage,
		Window:              &window,
		PullRequestURLModal: PullRequestURLModal,
		ProgressBar:         progressBar,
	}

	return criticWindow
}
