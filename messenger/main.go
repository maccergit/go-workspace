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
		ActualEmail{Target: "mcalisterjohnathon@gmail.com"},
		ActualEmail{Target: "john.mcalister@ricoh-usa.com"},
		ActualEmail{
			Target:       "example@example.com",
			SmtpOverride: "kendrick71@ethereal.email@smtp.ethereal.email",
			SmtpPass:     "P8fFs2Xy2ChMfGJ8Zs",
		},
		Email2SMS{
			ActualEmail: ActualEmail{Target: "4064103041"},
			Gateway:     "smtp.ethereal.mail",
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
