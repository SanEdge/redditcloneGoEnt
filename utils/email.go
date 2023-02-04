package utils

import (
	"fmt"
	"log"

	"github.com/renaldyhidayatt/redditgoent/dto/request"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(notificationEmail request.NotificationEmail) {
	message := gomail.NewMessage()
	message.SetHeader("From", "springreddit@email.com")
	message.SetHeader("To", notificationEmail.Recipient)
	message.SetHeader("Subject", notificationEmail.Subject)
	message.SetBody("text/plain", notificationEmail.Body)

	dialer := gomail.NewDialer(viper.GetString("SMTP_PROVIDER"), viper.GetInt("SMTP_PORT"), viper.GetString("SMTP_EMAIL"), viper.GetString("SMTP_PASSWORD"))
	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("Error occurred when sending mail: %v", err)
		return
	}
	fmt.Println("Activation email sent!")
}
