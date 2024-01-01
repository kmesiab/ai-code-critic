package ui

import (
	"fyne.io/fyne/v2"

	"github.com/kmesiab/ai-code-critic/ui/components"
)

type CriticWindow struct {
	Size                  fyne.Size // The size of the application window
	App                   *fyne.App
	Canvas                *fyne.Container                 // A box containing the ui components
	Window                *fyne.Window                    // The main application window
	ToolBar               *fyne.CanvasObject              // the top toolbar menu
	ReportPanel           *components.ReportPanel         // the left panel
	DiffPanel             *components.DiffPanel           // The right panel
	ProgressBar           *components.ProgressBar         // The progress bar
	PullRequestURLModal   *components.PullRequestURLModal // The API key modal
	IsAnalyzeButtonActive bool                            // flag to sprevent spamming of analyze button
}
