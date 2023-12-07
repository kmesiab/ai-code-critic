package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

func NewToolBar(
	menuButtonHandler critic.MenuButtonClickedEventHandler,
	fileOpenButtonHandler critic.FileOpenClickedEventHandler,
	analyzeButtonHandler critic.AnalyzeButtonClickedEventHandler,
) fyne.CanvasObject {
	toolbar := widget.NewToolbar(

		// Home button
		widget.NewToolbarAction(theme.MenuIcon(), menuButtonHandler),

		// File open button
		widget.NewToolbarAction(theme.FileTextIcon(), fileOpenButtonHandler),

		// Analyze button
		widget.NewToolbarAction(theme.ConfirmIcon(), analyzeButtonHandler),
	)

	return container.NewVBox(toolbar)
}
