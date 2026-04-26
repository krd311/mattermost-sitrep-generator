package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	var messageFilename string
	fmt.Print("Enter the filename of the messages JSON: ")
	fmt.Scanln(&messageFilename)

	messages, err := readMessages(messageFilename)
	if err != nil {
		log.Fatalf("Error reading messages: %v", err)
	} else {
		fmt.Printf("Successfully read %d messages from %s\n", len(messages), messageFilename)
	}

	fmt.Println("Generating SITREP...")
	sitRep, err := generateSitRep(messages)
	if err != nil {
		log.Fatalf("Error generating SITREP: %v", err)
	} else {
		fmt.Println("SITREP generated successfully:")
		fmt.Println(sitRep)
	}
}
