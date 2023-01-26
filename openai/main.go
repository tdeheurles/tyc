package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	gogpt "github.com/sashabaranov/go-gpt3"
)

const request = `Suggest four names for an animal that is a superhero.

Animal: Cat
Names: Captain Sharpclaw, Agent Fluffball, The Incredible Feline
Animal: Dog
Names: Ruff the Protector, Wonder Canine, Sir Barks-a-Lot
Animal: Elephant
Names:`

// documentation : https://beta.openai.com/docs/introduction

type Credentials struct {
	OpenaiToken string `json:"openai_token"`
}

func main() {
	credentials := getCredentials()
	createDirectory("results")
	result, shouldReturn := executeOpenAiCall(credentials)
	if shouldReturn {
		return
	}
	writeResultsToFile(result)
}

func writeResultsToFile(result []byte) {
	f, err := os.Create("results/" + getCurrentDataTime() + ".json")
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Write(result)
	defer f.Close()
}

func getCredentials() Credentials {
	file, err := os.Open("credentials.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	credentials, err := credentialsFromJSON(file)
	if err != nil {
		panic(err)
	}
	return credentials
}

func executeOpenAiCall(credentials Credentials) ([]byte, bool) {
	client := gogpt.NewClient(credentials.OpenaiToken)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		Temperature: 0.6,
		Prompt:      request,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		return nil, true
	}

	result, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	fmt.Println(string(result))
	return result, false
}

func credentialsFromJSON(file *os.File) (Credentials, error) {
	decoder := json.NewDecoder(file)
	credentials := Credentials{}
	err := decoder.Decode(&credentials)
	return credentials, err
}

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
}

func getCurrentDataTime() string {
	return time.Now().Format("2006-01-02_15:04:05")
}
