package internal

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
)

const (
	ApplicationName                = "AI Code Critic"
	PullRequestURLModalDefaultText = "Enter the github url for a pull request"
	MainCanvasHeight               = 720
	MainCanvasWidth                = 960
)

// Variables for the diff grid
var (
	greenColor = color.RGBA{R: 144, G: 238, B: 144, A: 255}
	redColor   = color.RGBA{R: 250, G: 128, B: 114, A: 255}
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

AI Code Critic is a tool that uses machine learning

to analyze your code and provide feedback on how to

improve it. It's like having a code reviewer in your

pocket!

![Logo](./assets/logo.png)

`

const WaitingForReportMarkdown = `
# Waiting for report... 🕐

Please be patient. Your report will appear here shortly.

![Logo](./assets/logo.png)
`
