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
		ActualEmail{SMTPServer: "smtp.gmail.com"},
		Email2SMS{
			ActualEmail: ActualEmail{SMTPServer: "smtp.provider.com"},
			Gateway:     "vtext.com",
		},
	}

	fmt.Printf("--- Broadcasting: %s ---\n", message)

	for _, recipent := range recipients {
		if err := recipent.Send(message); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}
	}

	fmt.Println("--- Broadcast Complete ---")
}
