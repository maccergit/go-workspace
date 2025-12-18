package main

import (
	"fmt"
	"os"
)

func main() {
	// If no argument is provided, we use a sensible default.
	message := "Default Broadcast: Hello from the Go Workspace!"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	recipients := []Notifier{
		MockEmail{Address: "boss@example.com"},
		MockEmail{Address: "hr@example.com"},
		MockSMS{Number: "555-0101"},
		MockSMS{Number: "555-0999"},
	}

	fmt.Printf("--- Broadcasting: %s ---\n", message)

	for _, r := range recipients {
		if err := r.Send(message); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}
	}

	fmt.Println("--- Broadcast Complete ---")
}
