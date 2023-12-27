package components

import (
	"fmt"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewToolBar(t *testing.T) {
	// setup
	testApp := test.NewApp()
	defer testApp.Quit()
	var pullRequestEventTriggered, analyzeEventTriggered bool

	// Mock event handlers
	pullRequestMenuItemClickedEventHandler := func() {
		pullRequestEventTriggered = true
	}

	analyzeButtonHandler := func() {
		analyzeEventTriggered = true
	}

	toolbarObject := NewToolBar(pullRequestMenuItemClickedEventHandler, analyzeButtonHandler)
	assert.NotNil(t, toolbarObject, "Toolbar should not be nil")

	window := test.NewWindow(toolbarObject)
	defer window.Close()

	toolbar := window.Canvas().Content().(*fyne.Container).Objects[0].(*widget.Toolbar)

	t.Run("Test toolbar actions handlers are activated", func(t *testing.T) {
		for _, item := range toolbar.Items {
			fmt.Printf("Item is %v", item)
			if action, ok := item.(*widget.ToolbarAction); ok {
				action.OnActivated()
			}
		}

		assert.True(t, pullRequestEventTriggered, "Pull request event should be triggered")
		assert.True(t, analyzeEventTriggered, "Analyze event should be triggered")
	})

}
