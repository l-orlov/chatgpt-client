package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey, gpt3.WithTimeout(2*time.Minute))

	const inputFile = "slim-fairy/in-out/input.txt"
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	msg := string(fileBytes)

	outputBuilder := strings.Builder{}

	err = client.ChatCompletionStream(ctx, gpt3.ChatCompletionRequest{
		Messages: []gpt3.ChatCompletionRequestMessage{
			{
				Role:    "user",
				Content: msg,
			},
		},
		MaxTokens:   3000,
		Temperature: gpt3.Float32Ptr(0.7),
	}, func(resp *gpt3.ChatCompletionStreamResponse) {
		outputBuilder.WriteString(resp.Choices[0].Delta.Content)
	})
	if err != nil {
		log.Fatalln(err)
	}
	output := strings.TrimSpace(outputBuilder.String())

	//err = client.CompletionStreamWithEngine(ctx, gpt3.GPT3Dot5Turbo, gpt3.CompletionRequest{
	//	Prompt: []string{
	//		msg,
	//	},
	//	MaxTokens:   gpt3.IntPtr(3000),
	//	Temperature: gpt3.Float32Ptr(0),
	//}, func(resp *gpt3.CompletionResponse) {
	//	outputBuilder.WriteString(resp.Choices[0].Text)
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//output := strings.TrimSpace(outputBuilder.String())

	const outputFile = "slim-fairy/in-out/output.txt"
	err = os.WriteFile(outputFile, []byte(output), os.ModePerm)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
}
