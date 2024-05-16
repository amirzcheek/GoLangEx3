package main

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChatCompletion(t *testing.T) {
	client := openai.NewClient("inmyhead")
	userInput := "Hello, how are you?"

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userInput,
				},
			},
		},
	)

	if err != nil {
		t.Fatalf("ChatCompletion error: %v", err)
	}

	if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
		t.Errorf("Expected a response from ChatGPT")
	}

	t.Logf("Response: %s", resp.Choices[0].Message.Content)
}

func TestHandleRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader("question=test"))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handleRequest)
	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.StatusCode)
	}
}
