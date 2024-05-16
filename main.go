package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

var (
	client      *openai.Client
	retryDelay  = 20 * time.Second
	maxRetries  = 3
	retryCount  = 0
	prevInput   string
	logger      *log.Logger
	filterWords = []string{"alcohol", "18+", "drugs"}
)

func init() {
	aikey := "inmyhead"
	client = openai.NewClient(aikey)

	logFile, err := os.OpenFile("request_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userInput := r.FormValue("question")
		if containsFilterWord(userInput, filterWords) {
			fmt.Fprint(w, "Declined: Your request was declined because your question is not suitable for children.")
			return
		}

		logger.Printf("Request: %s\n", userInput)
		userInput = prevInput + " " + userInput
		response, err := getChatResponse(userInput)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprint(w, response)
	} else {
		http.ServeFile(w, r, "index.html")
	}
}

func containsFilterWord(input string, words []string) bool {
	input = toLowerCase(input)
	for _, word := range words {
		strings.ToLower(input)
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}

func toLowerCase(str string) string {
	str = strings.ToLower(str)
	return str
}

func getChatResponse(input string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		if strings.Contains(err.Error(), "Rate limit reached") {
			retryCount++
			if retryCount > maxRetries {
				return "", fmt.Errorf("max retry limit reached")
			}
			time.Sleep(retryDelay)
			return getChatResponse(input)
		}
		return "", err
	}

	retryCount = 0

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}
	return "Sorry, I couldn't understand that.", nil
}
