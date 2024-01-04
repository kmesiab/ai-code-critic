package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

type ReportPanel struct {
	Size   fyne.Size
	Canvas *widget.RichText
}

func (p *ReportPanel) Resize() {
	p.Canvas.Resize(p.Size)
}

func (p *ReportPanel) SetDefaultText() {
	p.Canvas.ParseMarkdown(critic.IntroMarkdown)
}

func (p *ReportPanel) SetText(markdown string) {
	p.Canvas.ParseMarkdown(markdown)
}

func NewReportPanel(containerSize fyne.Size, text string) *ReportPanel {
	// Set it to half the width of the parent container
	newSize := fyne.NewSize(containerSize.Width/2, containerSize.Height)
	richText := widget.NewRichTextFromMarkdown(text)
	richText.Scroll = container.ScrollBoth
	richText.Wrapping = fyne.TextWrapWord
	richText.Resize(newSize)

	return &ReportPanel{Canvas: richText, Size: newSize}
}
