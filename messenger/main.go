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
		WebNotifier{URL: "https://webhook.site/bf0b1f13-807e-4fe6-9c9f-8e24c0673af6"},
		WebNotifier{URL: "https://ntfy.sh/johnmc_go_test_topic"},
	}

	fmt.Printf("--- Broadcasting: %s ---\n", message)

	for _, recipent := range recipients {
		if err := recipent.Send(message); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}
	}

	fmt.Println("--- Broadcast Complete ---")
}
