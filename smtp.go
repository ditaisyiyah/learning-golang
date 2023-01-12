package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	// smtp server configuration.
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

var (
	// sender account
	// ensure you disable the 2FA or set App password
	from     = "your_email"
	password = "your_email_password"

	// receipients
	to = []string{"receipient_email_1", "receipient_email_2"}

	// message
	message = "To:" +
		strings.Join(to, ";") +
		"\r\n" +
		"Subject: Greeting from Gophers!\r\n" +
		"\r\n" +
		"sstt! i love golang so muccchhh hahhaha "
)

func main() {
	// auth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))

	if err != nil {
		fmt.Println("Error happened: " + err.Error())
	}
}
