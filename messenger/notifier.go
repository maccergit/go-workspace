package main

// Notifier defines the "Contract" for any messaging service.
type Notifier interface {
	Send(message string) error
}
