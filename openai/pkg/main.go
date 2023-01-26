package lib

import (
	"fmt"
	"time"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func DoAnOpenAiQuery(request gogpt.CompletionRequest) {
	credentials := getCredentials()
	response := executeOpenAiCall(request, credentials)
	result := getResult(request, response)
	resultYaml := getResultAsYaml(result)

	fmt.Println(string(resultYaml))

	now := time.Now()
	directory := saveDirectory + "/" + now.Format("2006-01-02")
	createDirectory(directory)
	saveToFile(resultYaml, directory+"/"+now.Format("15:04:05")+"__"+credentials.User+".yaml")
}
