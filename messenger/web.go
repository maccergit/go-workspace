// messenger/web.go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Ensure WebNotifier satisfies the Notifier interface
var _ Notifier = (*WebNotifier)(nil)

type WebNotifier struct {
	URL string
}

// This is what we need to satisfy the Notifier interface
func (w WebNotifier) Send(msg string) error {
	resp, err := http.Post(w.URL, "text/plain", strings.NewReader(msg))
	if err != nil {
		return fmt.Errorf("web request failed: %w", err)
	}

	defer resp.Body.Close()

	// In HTTP, a "success" is anything in the 200-299 range.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned error status: %s", resp.Status)
	}

	fmt.Printf("[WEB] Successfully notified %s\n", w.URL)
	return nil
}
