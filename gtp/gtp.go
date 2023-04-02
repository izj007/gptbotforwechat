package gtp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/869413421/wechatbot/config"
)

const BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 请求体
//type ChatGPTResponseBody struct {
//	ID      string                   `json:"id"`
//	Object  string                   `json:"object"`
//	Created int                      `json:"created"`
//	Model   string                   `json:"model"`
//	Choices []map[string]interface{} `json:"choices"`
//	Usage   map[string]interface{}   `json:"usage"`
//}

type ChatGPTResponseBody struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int `json:"index"`
		Message      Message
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type ChoiceItem struct {
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTRequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatGPTRequestBody 响应体
//type ChatGPTRequestBody struct {
//	Model            string  `json:"model"`
//	Prompt           string  `json:"prompt"`
//	MaxTokens        int     `json:"max_tokens"`
//	Temperature      float32 `json:"temperature"`
//	TopP             int     `json:"top_p"`
//	FrequencyPenalty int     `json:"frequency_penalty"`
//	PresencePenalty  int     `json:"presence_penalty"`
//}

// Completions gtp文本模型回复
//curl https://api.openai.com/v1/completions
//-H "Content-Type: application/json"
//-H "Authorization: Bearer your chatGPT key"
//-d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	//	requestBody := ChatGPTRequestBody{
	//		Model:            "gpt-3.5-turbo",
	//		Prompt:           msg,
	//		MaxTokens:        2048,
	//		Temperature:      0.7,
	//		TopP:             1,
	//		FrequencyPenalty: 0,
	//		PresencePenalty:  0,
	//	}
	requestBody := ChatGPTRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: msg,
			},
		},
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"chat/completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}
	var reply string
	if len(gptResponseBody.Choices) > 0 {
		for _, v := range gptResponseBody.Choices {
			//reply = v["text"].(string)
			reply = v.Message.Content
			break
		}
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
