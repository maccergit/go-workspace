package main

import "fmt"

// These lines ensure at compile-time that our mocks satisfy the Notifier interface.
// If the interface changes, these lines will cause a compiler error here.
var _ Notifier = (*MockEmail)(nil)
var _ Notifier = (*MockSMS)(nil)

type MockEmail struct {
	Address string
}

func (m MockEmail) Send(msg string) error {
	fmt.Printf("[MOCK EMAIL] To: %s | Message: %s\n", m.Address, msg)
	return nil
}

type MockSMS struct {
	Number string
}

func (m MockSMS) Send(msg string) error {
	fmt.Printf("[MOCK SMS] To: %s | Message: %s\n", m.Number, msg)
	return nil
}
