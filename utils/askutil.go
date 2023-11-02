package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/tls-client/profiles"

	tls_client "github.com/bogdanfinn/tls-client"
)

type Data struct {
	Version string `json:"version"`
}

type Response struct {
	ID      string `json:"id"`
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

func GetWholeText(input string) (string, error) {
	if len(input) > 4000 {
		log.Fatalf("Input exceeds the input limit of 4000 characters")
	}

	resp, err := newRequest(input)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	code := resp.StatusCode

	if code >= 400 {
		return "", errors.New("status code 400")
	}

	scanner := bufio.NewScanner(resp.Body)

	fullText := ""

	for scanner.Scan() {
		mainText := getMainText(scanner.Text())
		fullText += mainText
	}

	return fullText, nil
}

func newRequest(input string) (*http.Response, error) {
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(120),
		tls_client.WithClientProfile(profiles.Chrome_103),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
	}...)
	if err != nil {
		return &http.Response{}, err
	}

	cleanInput, err := json.Marshal(input)
	if err != nil {
		return &http.Response{}, err
	}

	var reqBody = strings.NewReader(fmt.Sprintf(`{
		"frequency_penalty": 0,
		"messages": [
			{
				"content": %v,
				"role": "user"
			}
		],
		"model": "gpt-3.5-turbo",
		"presence_penalty": 0,
		"stream": true,
		"temperature": 1,
		"top_p": 1
	}
	`, string(cleanInput)))

	req, err := http.NewRequest("POST", "https://ai.fakeopen.com/v1/chat/completions", reqBody)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer pk-this-is-a-real-free-pool-token-for-everyone")

	return (client.Do(req))
}

func getMainText(line string) (mainText string) {
	var obj = "{}"
	if len(line) > 1 {
		obj = strings.Split(line, "data: ")[1]
	}

	var d Response
	if err := json.Unmarshal([]byte(obj), &d); err != nil {
		return ""
	}

	if d.Choices != nil {
		mainText = d.Choices[0].Delta.Content
		return mainText
	}
	return ""
}
