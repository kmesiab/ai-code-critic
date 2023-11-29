package internal

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
)

const (
	ApplicationName  = "AI Code Critic"
	MainCanvasHeight = 720
	MainCanvasWidth  = 960
)

// Variables for the diff grid
var (
	greenColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	redColor   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blackColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	whiteColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	RedTextGridStyle = &widget.CustomTextGridStyle{
		BGColor: redColor,
		FGColor: blackColor,
	}
	GreenTextGridStyle = &widget.CustomTextGridStyle{
		BGColor: greenColor,
		FGColor: blackColor,
	}
	BlackTextGridStyle = &widget.CustomTextGridStyle{
		BGColor: whiteColor,
		FGColor: blackColor,
	}
)

const IntroMarkdown = `
# Welcome to AI Code Critic! 🤖

![Logo](./assets/logo.png)

AI Code Critic is a tool that uses machine learning

to analyze your code and provide feedback on how to

improve it. It's like having a code reviewer in your

pocket!

`

const MoreInfoMarkdown = `
# Load a diff to get started 🚀

1. Click the "Open File" button in the top left corner. 📂

2. Select a file to analyze. 📄

3. Click the "Analyze" button. 🔍

4. Marvel over your new code review! 🎊

![Drag and drop](./assets/drag-and-drop.png)

`

const DragAndDropMarkdown = `
# Load a diff to get started
![Drag and drop](./assets/drag-and-drop.png)
`
