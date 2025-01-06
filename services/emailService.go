package services

import (
	"fmt"
	"net/smtp"
)

// SendOTPEmail sends an OTP to the user's email
func SendOTPEmail(email, otp string) error {
	// Set up the sender's email credentials (mock example, replace with actual credentials)
	sender := "your-email"
	password := "your-email-password"
	host := "smtp.gmail.com"
	port := "587"

	// Set up the recipient's email
	to := []string{email}

	// Set the email body
	subject := "Email Verification OTP"
	body := fmt.Sprintf("Your OTP for email verification is: %s", otp)

	// Set up the message
	message := []byte("Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		body)

	// Authentication
	auth := smtp.PlainAuth("", sender, password, host)

	// Send email
	err := smtp.SendMail(host+":"+port, auth, sender, to, message)
	if err != nil {
		return err
	}

	return nil
}
