# OpenAI API

## configure
Go to https://beta.openai.com/docs/introduction  
On upper right, click on `Personal` then `View API Keys` and `Create new secret key`  
Create a file `credentials.json` in this directory. Add your user name andsave your key like this:  
```json
{
  "user": "thibault",
  "openai_token": "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Do an example query
- [Install golang](https://go.dev/doc/install) on your machine  
- Open [`./query.go`](query.go) in vscode  
- Press `F5`  
- You should see the result printed as a `json`.  
- By design we decided to write the full API answer instead of just the text like in `chatGPT`.  

## Define your own query
You will see 2 variables `configuration` and `prompt`:
- `prompt` is the text you want to send to `openai`, just update it to run the query you want
- `configuration` is how to setup the model for the answer. It proposes several options. Have a look to [openai documentation](https://beta.openai.com/docs/introduction) for the details

## Save results
By design we decided to save results of each query into files.  
Each question you ask will be saved with its results in a file named by the datetime of execution.  
Ie: [`results/2023-01-26_12:00:04.json`](results/2023-01-26_12:00:04.json)  
