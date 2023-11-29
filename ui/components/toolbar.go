package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MenuButtonClickedEventHandler func()
type FileOpenClickedEventHandler func()
type AnalyzeButtonClickedEventHandler func()

func NewToolBar(
	menuButtonHandler MenuButtonClickedEventHandler,
	fileOpenButtonHandler FileOpenClickedEventHandler,
	analyzeButtonHandler AnalyzeButtonClickedEventHandler,
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
