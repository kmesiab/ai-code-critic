package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ReportPanel struct {
	Canvas *widget.RichText
	Size   fyne.Size
}

func NewReportPanel(containerSize fyne.Size, text string) *ReportPanel {
	return &ReportPanel{
		Size:   containerSize,
		Canvas: widget.NewRichTextFromMarkdown(text),
	}
}
