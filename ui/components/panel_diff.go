package components

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

type DiffPanel struct {
	Canvas   *container.Scroll // the scrollable parent
	TextGrid *widget.TextGrid  // the actual widget
	Size     fyne.Size
	MaxLines int
}

func NewDiffPanel(size fyne.Size, text string) *DiffPanel {
	newSize := fyne.NewSize(size.Width/.25, size.Height)

	grid := widget.NewTextGrid()
	grid.ShowLineNumbers = true

	scrollableParent := container.NewScroll(grid)

	panel := DiffPanel{
		Canvas:   scrollableParent,
		TextGrid: grid,
		Size:     newSize,
		MaxLines: 500,
	}

	return panel.SetDiffText(text)
}

func (grid *DiffPanel) SetDiffText(text string) *DiffPanel {
	lines := strings.Split(text, "\n")

	// limit the number of lines to 50
	if len(lines) > grid.MaxLines {
		lines = lines[:grid.MaxLines]
	}

	for i, line := range lines {

		var style *widget.CustomTextGridStyle

		if strings.HasPrefix(line, "+") {
			style = critic.GreenTextGridStyle
		} else if strings.HasPrefix(line, "-") {
			style = critic.RedTextGridStyle
		} else {
			style = critic.BlackTextGridStyle
		}

		textGridRow := widget.TextGridRow{Style: style}

		// Style the sentence, rune by rune
		for _, r := range line {
			textGridRow.Cells = append(
				textGridRow.Cells,
				widget.TextGridCell{Rune: r},
			)
		}

		grid.TextGrid.SetRow(i, textGridRow)

	}

	return grid
}
