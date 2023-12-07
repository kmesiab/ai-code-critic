package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

func NewToolBar(
	fileOpenButtonHandler critic.FileOpenClickedEventHandler,
	analyzeButtonHandler critic.AnalyzeButtonClickedEventHandler,
) fyne.CanvasObject {
	toolbar := widget.NewToolbar(

		// File open button
		widget.NewToolbarAction(theme.ContentAddIcon(), fileOpenButtonHandler),

		// Analyze button
		widget.NewToolbarAction(theme.MediaPlayIcon(), analyzeButtonHandler),
	)

	return container.NewVBox(toolbar)
}
