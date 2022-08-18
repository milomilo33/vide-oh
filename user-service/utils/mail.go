package utils

import (
	"fmt"
	"net/smtp"
)

func SendBlockedMail(email string) {

	// Sender data.
	from := "vide.oh@smtp.com"
	// password := "strongpw#1"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "localhost"
	smtpPort := "1025"

	// Message.
	message := []byte("Your account has been blocked.")

	// Authentication.
	// auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
