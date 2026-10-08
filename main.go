package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	// Retrieve SMTP server details from environment variables.
	// Bu kısım, SMTP sunucu bilgilerini ortam değişkenlerinden alır.
	// Güvenlik için hassas bilgileri doğrudan kodda tutmaktan kaçının.
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	senderEmail := os.Getenv("SENDER_EMAIL")
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")

	// Check if all required environment variables are set.
	if smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" || senderEmail == "" || recipientEmail == "" {
		log.Fatal("Error: Please set SMTP_HOST, SMTP_PORT, SMTP_USERNAME, SMTP_PASSWORD, SENDER_EMAIL, and RECIPIENT_EMAIL environment variables.")
	}

	// Authentication for the SMTP server.
	// SMTP sunucusu için kimlik doğrulama bilgileri.
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Construct the email message including headers.
	// E-posta mesajının başlıkları ve içeriği oluşturuluyor.
	// "From", "To", "Subject" başlıkları mesajın bir parçası olmalıdır.
	message := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Go ile Hoş Geldin E-postası Denemesi\r\n"+ // Turkish for "Go Welcome Email Test"
		"MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"+
		"\r\n"+ // Empty line separates headers from body
		"Merhaba,\r\n\r\nBu, Go kullanarak gönderilen bir hoş geldin e-postası denemesidir.\r\n\r\nSaygılarımızla,\r\nGo E-posta Servisi", // Turkish for "Hello, This is a welcome email test sent using Go. Regards, Go Email Service"
		senderEmail, recipientEmail)

	msg := []byte(message)

	// Send the email.
	// smtp.SendMail fonksiyonu, belirtilen sunucuya bağlanır, kimlik doğrulaması yapar ve e-postayı gönderir.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, msg)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	fmt.Printf("From: %s\nTo: %s\n", senderEmail, recipientEmail)
}
