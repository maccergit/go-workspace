package main

import "fmt"

var _ Notifier = (*ActualEmail)(nil)
var _ Notifier = (*Email2SMS)(nil)

type ActualEmail struct {
	SMTPServer string
}

func (e ActualEmail) Send(msg string) error {
	fmt.Printf("[REAL EMAIL] Sending via %s: %s\n", e.SMTPServer, msg)
	return nil
}

type Email2SMS struct {
	ActualEmail // no field name, just type - signatures are promoted
	Gateway     string
}
