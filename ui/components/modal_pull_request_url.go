package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
)

type PullRequestURLModal struct {
	Form      *dialog.FormDialog
	TextEntry *widget.Entry
	GPTModel  *widget.SelectEntry
}

func NewPullRequestURLModal(
	size fyne.Size,
	defaultText string,
	parentWindow *fyne.Window,
	onSubmitHandler critic.OnSubmitButtonClickedEvent,
) *PullRequestURLModal {
	textEntry := widget.NewEntry()
	textEntryFormItem := widget.NewFormItem("Github PR URL", textEntry)
	gptModelEntry := widget.NewSelectEntry(critic.SupportedGPTModels)
	gptModelEntryFormItem := widget.NewFormItem("GPT Model", gptModelEntry)

	formItems := []*widget.FormItem{textEntryFormItem, gptModelEntryFormItem}

	f := dialog.NewForm(
		defaultText,
		"Submit",
		"Cancel",
		formItems,
		onSubmitHandler,
		*parentWindow,
	)

	newSize := fyne.NewSize(size.Width/.7, f.MinSize().Height)
	f.Resize(newSize)

	return &PullRequestURLModal{Form: f, TextEntry: textEntry, GPTModel: gptModelEntry}
}
