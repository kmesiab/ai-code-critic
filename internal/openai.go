package internal

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func GetAPIKey() (string, error) {
	cfg, err := GetConfig()
	if err != nil {
		return "", err
	}

	if cfg.OpenAIAPIKey == "" {
		return "", fmt.Errorf("API key not set in environment variables")
	}

	return cfg.OpenAIAPIKey, nil
}

func GetCodeReviewFromAPI(diff, gptModel string) (string, error) {
	apiKey, err := GetAPIKey()
	if err != nil {
		return "", err
	}

	prompt := `You are an experienced software developer conducting a code review on a Git diff. Your expertise spans 
	various programming languages and development best practices. Please review the attached Git diff with the 
	following considerations in mind:

1. **Technical Accuracy**: Identify any bugs, coding errors, or security vulnerabilities.
2. **Best Practices**: Evaluate adherence to language-specific best practices, including code style and patterns.
3. **Performance and Scalability**: Highlight any performance issues and assess the code's scalability.
4. **Readability and Clarity**: Assess the code's readability, including its structure and commenting.
5. **Maintainability**: Consider the ease of future modifications and support.
6. **Testability**: Evaluate the test coverage and quality of tests.
7. **Contextual Fit**: Judge how the changes fit within the broader project scope and goals.

Provide actionable feedback, suggesting improvements and alternatives where applicable.  Include code samples in code 
blocks. Your review should be empathetic and constructive, focusing on helping the author improve the code. 
Format your review in markdown, ensuring readability with line wrapping before 60 characters.

In your review, consider the impact of your feedback on team dynamics and the development process. Aim for a 
balance between technical rigor and fostering a positive and collaborative team environment.

Output the review in markdown format

### Git Diff:` + "```\n%s\n```"

	fullPrompt := fmt.Sprintf(prompt, diff)

	// OpenAI
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: gptModel,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fullPrompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
