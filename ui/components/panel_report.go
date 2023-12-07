package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ReportPanel struct {
	Canvas *widget.RichText
}

func (panel *ReportPanel) Resize(size *fyne.Size) *ReportPanel {
	panel.Canvas.Resize(*size)
	return panel
}

func NewReportPanel(containerSize fyne.Size, text string) *ReportPanel {

	// Set it to half the width of the parent container
	newSize := fyne.NewSize(containerSize.Width/2, containerSize.Height)
	richText := widget.NewRichTextFromMarkdown(text)
	richText.Resize(newSize)

	return &ReportPanel{Canvas: richText}
}
