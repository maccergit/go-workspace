package main

import (
	"fmt"
	"os"
)

var _ Notifier = (*ActualEmail)(nil)
var _ Notifier = (*Email2SMS)(nil)

type ActualEmail struct {
	SMTPServer string
}

func (e ActualEmail) Send(msg string) error {
	fmt.Printf("[REAL EMAIL] Sending via %s: %s\n", e.SMTPServer, msg)

	apiKey := os.Getenv("SMTP2GO_API_KEY")

	if apiKey == "" {
		return fmt.Errorf("SMTP2GO_API_KEY is not set")
	}

	return nil
}

type Email2SMS struct {
	ActualEmail // no field name, just type - signatures are promoted
	Gateway     string
}
