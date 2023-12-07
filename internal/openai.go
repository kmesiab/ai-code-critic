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

func GetCodeReviewFromAPI(diff string) (string, error) {
	apiKey, err := GetAPIKey()

	if err != nil {
		return "", err
	}

	prompt := `
	You are a helpful software developer. 
	Attached below is a diff of a pull request.
	You will review the diff for bugs and common mistakes.
	You will review the diff for best practices according to its language.
	You will review the diff for security vulnerabilities.
	You will review the diff for performance issues.
	You will review the diff for readability.
	You will review the diff for maintainability.
	You will review the diff for testability.
	You will review the diff for scalability.
	You will review the diff for extensibility.
	You will review the diff for reusability.
	You will review the diff for modularity.
	You will review the diff for simplicity.
	You will produce a concise code review and output it in markdown.
	Break all lines before 60 characters.
	
	Diff:
	%s
`

	fullPrompt := fmt.Sprintf(prompt, diff)

	// OpenAI
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
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
