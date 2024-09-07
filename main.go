package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	fmt.Println("Hello")

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_FROM_PASSWORD")
	to := []string{os.Getenv("EMAIL_TO")}
	// url := os.Getenv("URL")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte("This is a test message.")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent")
}
