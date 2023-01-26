package lib

import (
	"encoding/json"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
	"gopkg.in/yaml.v3"
)

const saveDirectory = "saves"

type Result struct {
	Query    string                   `json:"query" yaml:"query"`
	Response gogpt.CompletionResponse `json:"response" yaml:"response"`
}

func saveToFile(content []byte, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(content)
}

func credentialsFromFile(file *os.File) (Credentials, error) {
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

func getResultAsYaml(result Result) []byte {
	resultYaml, err := yaml.Marshal(result)
	if err != nil {
		panic(err)
	}
	return resultYaml
}

func getResult(request gogpt.CompletionRequest, response gogpt.CompletionResponse) Result {
	result := Result{
		Query:    request.Prompt,
		Response: response,
	}
	return result
}
