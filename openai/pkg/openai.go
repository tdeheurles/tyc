package lib

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func executeOpenAiCall(request gogpt.CompletionRequest, credentials Credentials) gogpt.CompletionResponse {
	client := gogpt.NewClient(credentials.OpenaiToken)
	ctx := context.Background()
	resp, err := client.CreateCompletion(ctx, request)
	if err != nil {
		panic(err)
	}
	return resp
}
