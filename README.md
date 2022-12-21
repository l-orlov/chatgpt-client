# chatgpt-client
Golang client for for ChatGPI (<https://chat.openai.com>) using GPT

## Request for API from OpenAI
Request your OpenAPI key in [https://beta.openai.com/account/api-keys](https://beta.openai.com/account/api-keys)

## Configure project
- install go (<https://go.dev/doc/install>)
- run command in terminal: `go mod tidy`
- set your OpenAPI key in `.env` file: `API_KEY=your_key`

## Get list of libraries that are used in the code
- write python code snippet in file `input_with_code.txt`
- run go program with command in terminal: `go run code-libraries/main.go`
- see libraries list in file `output.txt`
