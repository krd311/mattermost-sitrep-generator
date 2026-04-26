package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		// Check for .env file for API key and error check
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, relying on system environment variables")
		}*/

	// take in input of filename to allow for different message files
	var messageFilename string

	// can't print and take input at the same time; print msg then take input
	fmt.Print("Enter the filename of the messages JSON: ")
	fmt.Scanln(&messageFilename)

	// read messages from json file
	// messages = the messages returned from readMessages, err != nil if there was an error
	messages, err := readMessages(fmt.Sprintf(`conversations\%v`, messageFilename))
	if err != nil {
		log.Fatalf("Error reading messages: %v", err)
	} else {
		fmt.Printf("Successfully read %d messages from %s\n", len(messages), messageFilename)
	}

	// generate SITREP from messages
	fmt.Println("Generating SITREP...")
	// sitRep == string representing response from openAI call, err != nil if there was an error
	sitRep, err := generateSitRep(messages)
	if err != nil {
		log.Fatalf("Error generating SITREP: %v", err)
	} else {
		fmt.Println("SITREP generated successfully:")
		fmt.Println(sitRep)
	}
}
