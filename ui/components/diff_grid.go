package components

import (
	"strings"

	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

type DiffGrid struct {
	widget.TextGrid
}

func NewDiffGrid(text string) *DiffGrid {
	grid := &DiffGrid{}
	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	return grid.SetText(text)
}

func (grid *DiffGrid) SetText(text string) *DiffGrid {
	lines := strings.Split(text, "\n")
	var rows []widget.TextGridRow

	for _, line := range lines {
		var style *widget.CustomTextGridStyle

		if strings.HasPrefix(line, "+") {
			style = critic.GreenTextGridStyle
		} else if strings.HasPrefix(line, "-") {
			style = critic.RedTextGridStyle
		} else {
			style = critic.BlackTextGridStyle
		}

		textGridRow := widget.TextGridRow{
			Style: style,
		}

		// Style the sentence, rune by rune
		for _, r := range line {
			textGridRow.Cells = append(
				textGridRow.Cells,
				widget.TextGridCell{Rune: r, Style: style},
			)
		}

		rows = append(rows, textGridRow)

	}

	grid.Rows = rows
	grid.Refresh()
	return grid
}
