package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {

	loadErr := godotenv.Load(".env")
	if loadErr != nil {
		log.Fatalf("Error loading .env file")
	}
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")
	url := os.Getenv("URL")

	response, requestErr := http.Get(url)
	if requestErr != nil {
		log.Fatalf(requestErr.Error())
	}
	defer response.Body.Close()

	body, readAllErr := io.ReadAll(response.Body)

	if readAllErr != nil {
		log.Fatalf(readAllErr.Error())
	}

	previousBody, previousBodyError := os.ReadFile("doc.html")

	writeErr := os.WriteFile("doc.html", body, 0644)

	if writeErr != nil {
		log.Fatalf(writeErr.Error())
	}
	if previousBodyError != nil {
		log.Fatalf(previousBodyError.Error())
	}
	if string(previousBody) == string(body) {
		log.Print("Bodies are the same, exiting")
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("Le site %s a recu une mise a jour", url))
	m.SetBody("text/plain", fmt.Sprintf("Va vite sur %s !", url))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Error sending mail to " + to)
	}
	fmt.Println("Email sent")
}
