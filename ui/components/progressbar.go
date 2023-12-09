package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ProgressBar struct {
	Size   fyne.Size
	Canvas *widget.ProgressBarInfinite
}

func NewProgressBar(containerSize fyne.Size) *ProgressBar {
	newSize := fyne.NewSize(containerSize.Width, containerSize.Height)
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Stop()

	return &ProgressBar{Canvas: progressBar, Size: newSize}
}

func (p *ProgressBar) Resize(newSize fyne.Size) {
	p.Canvas.Resize(newSize)
}

func (p *ProgressBar) IsRunning() bool {
	return p.Canvas.Running()
}

func (p *ProgressBar) StartProgressBar() {
	if !p.Canvas.Running() {
		p.Canvas.Start()
		p.Canvas.Show()
	}
}

func (p *ProgressBar) StopProgressBar() {
	if p.Canvas.Running() {
		p.Canvas.Stop()
		p.Canvas.Hide()
	}
}
