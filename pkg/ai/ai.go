package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GenerateReadme(description string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY 환경 변수가 설정되지 않았습니다")
	}
  
	client := openai.NewClient(apiKey)
	
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4, // 또는 GPT3Dot5Turbo
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: fmt.Sprintf("Generate a professional README.md for the following Go project:\n\n%s", description),
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}