package internal

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
	"github.com/sashabaranov/go-openai"
)

const (
	ApplicationName                = "AI Code Critic"
	PullRequestURLModalDefaultText = "Enter the GitHub URL for a pull request:"
	MainCanvasHeight               = 600
	MainCanvasWidth                = 800
)

var SupportedGPTModels = []string{
	openai.GPT432K0613,
	openai.GPT432K0314,
	openai.GPT432K,
	openai.GPT40613,
	openai.GPT40314,
	openai.GPT4TurboPreview,
	openai.GPT4VisionPreview,
	openai.GPT4,
	openai.GPT3Dot5Turbo1106,
	openai.GPT3Dot5Turbo0613,
	openai.GPT3Dot5Turbo0301,
	openai.GPT3Dot5Turbo16K,
	openai.GPT3Dot5Turbo16K0613,
	openai.GPT3Dot5Turbo,
	openai.GPT3Dot5TurboInstruct,
	openai.GPT3Davinci,
	openai.GPT3Davinci002,
	openai.GPT3Curie,
	openai.GPT3Curie002,
	openai.GPT3Ada,
	openai.GPT3Ada002,
	openai.GPT3Babbage,
	openai.GPT3Babbage002,
}

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
