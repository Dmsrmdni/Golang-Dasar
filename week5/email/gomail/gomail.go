package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("failed to convert SMTP_PORT to integer: %v", err)
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpEmail)
	mailer.SetHeader("To", "dimasramdani220@gmail.com")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("/var/www/Golang/Belajar-golang-dasar/week5/email/gomail/test.png")

	dialer := gomail.NewDialer(
		smtpHost,
		smtpPort,
		smtpEmail,
		smtpPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
