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

	diffPanel := components.NewDiffPanel(canvasSize, "")
	reportPanel := components.NewReportPanel(canvasSize, critic.IntroMarkdown)

	// The toolbar sits atop the horizontal container
	toolbar := components.NewToolBar(
		fileOpenButtonClickedHandler,
		analyzeButtonClickedHandler,
	)

	progressBar := components.NewProgressBar(canvasSize)

	// Lay out all the panels
	horizontalContainer := container.NewHBox(reportPanel.Canvas, diffPanel.Canvas)
	fullCanvas := container.NewVBox(toolbar, horizontalContainer, progressBar.Canvas)

	// Create a main window and set the canvas as its content
	window := app.NewWindow(critic.ApplicationName)
	window.SetContent(fullCanvas)
	window.Resize(canvasSize)
	window.SetFixedSize(true)

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
		Canvas:              fullCanvas,
		Window:              &window,
		PullRequestURLModal: PullRequestURLModal,
		ProgressBar:         progressBar,
	}

	return criticWindow
}
