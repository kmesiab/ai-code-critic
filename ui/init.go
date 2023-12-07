package ui

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	critic "github.com/kmesiab/ai-code-critic/internal"
	"github.com/kmesiab/ai-code-critic/ui/components"
)

type CriticWindow struct {
	App                 *fyne.App
	Size                fyne.Size                       // The size of the application window
	ReportPanel         *components.ReportPanel         // the left panel
	DiffPanel           *components.DiffPanel           // The right panel
	ToolBar             *fyne.CanvasObject              // the top toolbar menu
	CenterDivider       *widget.Separator               // A separator between the two panels
	Canvas              *fyne.Container                 // A vertical box containing the ui components
	Window              *fyne.Window                    // The main application window
	PullRequestURLModal *components.PullRequestURLModal // The API key modal
}

var criticWindow *CriticWindow

func Initialize(app fyne.App) *CriticWindow {
	canvasSize := fyne.NewSize(critic.MainCanvasWidth, critic.MainCanvasHeight)

	// Left and right rich text panels and a center divider
	left := components.NewReportPanel(canvasSize, critic.IntroMarkdown)
	right := components.NewDiffPanel(canvasSize, LoadSampleDiffString())
	center := widget.NewSeparator()

	// Three panels in a horizontal container
	horizontalContainer := container.NewHBox(left.Canvas, center, right.Canvas)

	// The toolbar sits atop the horizontal container
	toolbar := components.NewToolBar(
		onMenuButtonClickedHandler,
		onFileOpenButtonClickedHandler,
		onAnalyzeButtonClickedHandler,
	)

	// All then panels laid out in a vertical container
	fullCanvas := container.NewVBox(toolbar, horizontalContainer)

	// Create a main window and set the canvas as its content
	window := app.NewWindow(critic.ApplicationName)
	window.SetContent(fullCanvas)
	window.Resize(canvasSize)

	PullRequestURLModal := components.NewPullRequestURLModal(
		critic.PullRequestURLModalDefaultText, &window,
		onAPIKeySubmitButtonClickedHandler,
	)

	criticWindow = &CriticWindow{
		App:                 &app,
		Size:                canvasSize,
		ReportPanel:         left,
		DiffPanel:           right,
		ToolBar:             &toolbar,
		CenterDivider:       center,
		Canvas:              fullCanvas,
		Window:              &window,
		PullRequestURLModal: PullRequestURLModal,
	}

	return criticWindow
}

func LoadSampleDiffString() string {
	diffBytes, err := os.ReadFile("./assets/diff.txt")
	if err != nil {
		log.Println(err)
		return critic.MoreInfoMarkdown
	}

	return string(diffBytes)
}

func onAPIKeySubmitButtonClickedHandler(canceled bool) {
	input := criticWindow.PullRequestURLModal.TextEntry.Text
	log.Printf("Menu button clicked: %s", input)
}

func onMenuButtonClickedHandler() {
}

func onFileOpenButtonClickedHandler() {
	criticWindow.PullRequestURLModal.Form.Show()
}

func onAnalyzeButtonClickedHandler() {
	criticWindow.DiffPanel.SetText(LoadSampleDiffString())
}
