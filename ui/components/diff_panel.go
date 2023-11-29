package components

import (
	"fyne.io/fyne/v2"
)

type DiffPanel struct {
	Canvas *DiffGrid
	Size   fyne.Size
}

func NewDiffPanel(size fyne.Size, text string) *DiffPanel {

	grid := NewDiffGrid(text)
	grid.ShowLineNumbers = true

	grid.Resize(size)

	return &DiffPanel{
		Canvas: grid,
		Size:   size,
	}
}

func (panel *DiffPanel) SetText(text string) *DiffPanel {
	panel.Canvas.SetText(text)

	return panel
}
