package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func generateSitRep(messages []Message) (string, error) {
	// grab environment variable for openAI API key; return error if not set
	// Check for .env file for API key and error check
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	// openAI API call
	// msgBytes == messages turned into json
	msgBytes, err := json.Marshal(messages)
	if err != nil {
		return "", fmt.Errorf("failed to marshal messages: %v", err)
	}

	// prompt woo hoo
	// can be optimized to reduce token usage as well as return more consise SITREPs, but this is a good starting point
	// when requirements are clear, i would refine this
	// we would probably want to use a more powerful model if needed as well
	// sprintf == string formatting
	prompt := fmt.Sprintf(`You are an assistant that is tasked with taking a list of messages and generating a SITREP.
	This SITREP will be used in a military context. Make the SITREP concise while keeping all important information. Omit what is not important. Produce a timeline and brief suggestions for action. Here are the messages: %s`, string(msgBytes))

	requestBody := map[string]interface{}{
		"model": "gpt-4.1-mini",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	// convert request body to json, return error if it fails
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// create new POST request to openAI API, set headers, and execute request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make API call: %v", err)
	}

	// defer is weird.
	// it means to execute the function after the current function (generateSitRep) finishes executing
	defer response.Body.Close()

	// error checking
	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("API call failed with status %d: %s", response.StatusCode, string(bodyBytes))
	}

	// read response body, return error if it fails
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// create result variable representing response from API call
	var result result

	// if there is an error storing the response into result, return error
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	// return raw string containing message content
	return result.Choices[0].Message.Content, nil
}
