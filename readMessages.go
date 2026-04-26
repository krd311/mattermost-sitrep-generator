package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readMessages(filename string) ([]Message, error) {

	// get array of messages from input json file
	var messages []Message

	// data == raw bytes from the file, err != nil if there was an error reading the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// unmarshal... means to turn something from raw json to a go-readable var
	// store that go-readable var in messages
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return messages, nil
}
