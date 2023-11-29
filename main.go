//go:generate ai-code-critic bundle bundled.go assets

package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/kmesiab/ai-code-critic/ui"
)

func main() {
	application := app.New()
	window := ui.Initialize(application).Window
	(*window).ShowAndRun()
}
