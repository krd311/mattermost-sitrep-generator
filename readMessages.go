package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readMessages(filename string) ([]Message, error) {

	// Get array of messages from input json file
	var messages []Message
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return messages, nil
}
