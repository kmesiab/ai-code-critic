package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

func NewToolBar(
	pullRequestMenuItemClickedEventHandler critic.PullRequestMenuItemClickedEventHandler,
	analyzeButtonHandler critic.AnalyzeButtonClickedEventHandler,
) fyne.CanvasObject {
	toolbar := widget.NewToolbar(

		// Open pull request button
		widget.NewToolbarAction(theme.ContentAddIcon(), pullRequestMenuItemClickedEventHandler),

		// Analyze diff button
		widget.NewToolbarAction(theme.MediaPlayIcon(), analyzeButtonHandler),
	)

	return container.NewVBox(toolbar)
}
