package ui

import (
	"fyne.io/fyne/v2"

	"github.com/kmesiab/ai-code-critic/ui/components"
)

type CriticWindow struct {
	App                 *fyne.App
	Size                fyne.Size                       // The size of the application window
	ReportPanel         *components.ReportPanel         // the left panel
	DiffPanel           *components.DiffPanel           // The right panel
	ToolBar             *fyne.CanvasObject              // the top toolbar menu
	Canvas              *fyne.Container                 // A vertical box containing the ui components
	Window              *fyne.Window                    // The main application window
	PullRequestURLModal *components.PullRequestURLModal // The API key modal
}
