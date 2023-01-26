package lib

import "os"

type Credentials struct {
	User        string `json:"user" yaml:"user"`
	OpenaiToken string `json:"openai_token" yaml:"openai_token"`
}

func getCredentials() Credentials {
	file, err := os.Open("credentials.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	credentials, err := credentialsFromFile(file)
	if err != nil {
		panic(err)
	}
	return credentials
}
