package components

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestModalPullRequestURL(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()

	testWindow := testApp.NewWindow("Test Window")
	defer testWindow.Close()

	// Call the function to test
	modal := NewPullRequestURLModal(fyne.NewSize(600, 400), "Default Text", &testWindow, nil)

	// Check if modal is correctly initialized
	assert.NotNil(t, modal, "Modal should not be nil")
	assert.NotNil(t, modal.Form, "Form dialog should not be nil")
	assert.NotNil(t, modal.TextEntry, "Text entry should not be nil")
	assert.NotNil(t, modal.GPTModel, "GPT model select entry should not be nil")

	// Interact with the form
	test.Type(modal.TextEntry, "https://github.com/example/pr")
	test.Type(modal.GPTModel, "GPT-3")

	// Assert the state of UI elements after interactions
	assert.Equal(t, "https://github.com/example/pr", modal.TextEntry.Text, "Text entry should contain the typed URL")
	assert.Equal(t, "GPT-3", modal.GPTModel.Text, "GPT model select entry should contain the selected model")
}
