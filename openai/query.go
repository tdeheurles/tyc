package main

import (
	gogpt "github.com/sashabaranov/go-gpt3"
)

// =====================================================================================
// documentation : https://beta.openai.com/docs/introduction
var req = gogpt.CompletionRequest{
	Model:       gogpt.GPT3TextDavinci003,
	Temperature: 0.6,
	Prompt:      request,
}

const request = `Suggest four names for an animal that is a superhero.

Animal: Cat
Names: Captain Sharpclaw, Agent Fluffball, The Incredible Feline
Animal: Dog
Names: Ruff the Protector, Wonder Canine, Sir Barks-a-Lot
Animal: Elephant
Names:`

// =====================================================================================

func main() {
	doAnOpenAiQuery()
}
