/*
Simple program to show some unique Go elements :
- package and imports - note that elements with a first Capital letter are public
(this is a Go requirement - not simply a convention)
- mutiple return values - last value is typically an error status, as Go does not
have exceptions.
- "if with initialization" - common idiom used to check error condition after a call.
- Automatic Semicolon Insertion (ASI)
- Type inference with := operator.
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// extractNameFromArgs validates the presence of an argument and returns the
// raw name string or an error if the argument is missing.
func extractNameFromArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("missing command-line argument.\nUsage: go run greetings.go <name>")
	}
	return args[1], nil
}

// generateMessage validates the name content and returns the formatted greeting or an error.
func generateMessage(name string) (string, error) {
	cleanedName := strings.TrimSpace(name) 
	if cleanedName == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", cleanedName)
	return message, nil
}

// displayGreeting calls the business logic and handles printing the final result or error.
func displayGreeting(rawName string) {
	if message, err := generateMessage(rawName); err != nil {
		fmt.Printf("Error: Could not generate greeting: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(message)
	}
}

// main is the entry point of the executable program. It handles I/O validation.
func main() {
	if rawName, err := extractNameFromArgs(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		displayGreeting(rawName)
	}
}