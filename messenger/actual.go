package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var _ Notifier = (*ActualEmail)(nil)
var _ Notifier = (*Email2SMS)(nil)

type ActualEmail struct {
	SmtpOverride string // format: user@host:port (or user@domain@host:port)
	SmtpPass     string // ignored unless SmtpOverride is used
	Target       string
}

func parseSmtpConfig(smtpConfig string, smtpPass string) (user string, hostConfig string, pass string, err error) {
	if smtpConfig == "" {
		smtpConfig = os.Getenv("SMTP_ACCOUNT")
		smtpPass = os.Getenv("SMTP_PASS")
	}

	if smtpConfig == "" {
		return "", "", "", fmt.Errorf("configuration error: no SMTP config found (struct or env)")
	}

	if smtpPass == "" {
		return "", "", "", fmt.Errorf("auth error: no SMTP password found for '%s'", smtpConfig)
	}

	lastAtIndex := strings.LastIndex(smtpConfig, "@")

	if lastAtIndex == -1 {
		return "", "", "", fmt.Errorf("invalid config format: %s (missing @)", smtpConfig)
	}

	return smtpConfig[:lastAtIndex], smtpConfig[lastAtIndex+1:], smtpPass, nil
}

// Send satisfies the Notifier interface using the internal Target
func (e ActualEmail) Send(msg string) error {
	return e.SendTo(e.Target, msg)
}

// SendTo allows sending to a specific address using the struct's SMTP config
func (e ActualEmail) SendTo(target string, msg string) error {
	user, host, pass, err := parseSmtpConfig(e.SmtpOverride, e.SmtpPass)
	if err != nil {
		return err
	}

	hostParts := strings.Split(host, ":")
	server := hostParts[0]
	port := "587"
	if len(hostParts) > 1 {
		port = hostParts[1]
	}

	auth := smtp.PlainAuth("", user, pass, server)
	addr := fmt.Sprintf("%s:%s", server, port)
	header := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Go Messenger Alert\r\n\r\n", user, target)
	body := []byte(header + msg)

	if err := smtp.SendMail(addr, auth, user, []string{target}, body); err != nil {
		return fmt.Errorf("delivery failed to %s: %w", target, err)
	}

	fmt.Printf("[EMAIL] Successfully sent to %s\n", target)
	return nil
}

type Email2SMS struct {
	ActualEmail // no field name, just type - signatures are promoted
	Gateway     string
}

func (s Email2SMS) Send(msg string) error {
	// Construct the gateway address locally
	fullTarget := fmt.Sprintf("%s@%s", s.ActualEmail.Target, s.Gateway)

	// Delegate the work to the embedded SendTo method
	return s.ActualEmail.SendTo(fullTarget, msg)
}
