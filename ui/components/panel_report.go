package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ReportPanel struct {
	Size   fyne.Size
	Canvas *widget.RichText
}

func NewReportPanel(containerSize fyne.Size, text string) *ReportPanel {

	// Set it to half the width of the parent container
	newSize := fyne.NewSize(containerSize.Width/2, containerSize.Height)
	richText := widget.NewRichTextFromMarkdown(text)
	richText.Scroll = container.ScrollBoth
	richText.Resize(newSize)

	return &ReportPanel{Canvas: richText, Size: newSize}
}
