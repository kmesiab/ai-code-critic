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

	// Set it to half the width of the parent container
	newSize := fyne.NewSize(containerSize.Width*.3, containerSize.Height)
	progressBar := widget.NewProgressBarInfinite()
	progressBar.Stop()
	progressBar.Resize(newSize)

	return &ProgressBar{Canvas: progressBar, Size: newSize}
}
