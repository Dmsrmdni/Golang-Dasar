package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPort := os.Getenv("SMTP_PORT")

	// Alamat email penerima
	to := []string{"dimasramdani220@gmail.com"}

	// Pesan email
	subject := "Test Email"
	body := "Test Email"

	// Membuat pesan email
	message := "From: " + smtpEmail + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	// Mengirim email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpEmail, to, []byte(message))
	if err != nil {
		log.Fatal("Failed to send mail:", err)
	} else {
		log.Println("Email sent successfully!")
	}
}
