package main

// First, let's add the necessary packages to send emails and integrate with Slack:

import (
	"context"
	"fmt"
	"log"

	sg "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/slack-go/slack"
)

func sendEmailNotification(itemName string, daysBeforeExpiration int) {
	from := mail.NewEmail("Sender Name", "sender@example.com")
	subject := fmt.Sprintf("Expiration Reminder: %s", itemName)
	to := mail.NewEmail("Recipient Name", "recipient@example.com")
	plainTextContent := fmt.Sprintf("The item %s is expiring in %d days.", itemName, daysBeforeExpiration)
	htmlContent := fmt.Sprintf(`<p>The item **%s** is expiring in %d days.</p>`, itemName, daysBeforeExpiration)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sg.NewSendClient("YOUR_SENDGRID_API_KEY")
	_, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send email notification:", err)
	}
}

// To send a Slack notification
func SendSlackNotification(itemName string, daysBeforeExpiration int) {
	api := slack.New("YOUR_SLACK_API_TOKEN")
	channelID, _, err := api.PostMessageContext(
		context.TODO(),
		"YOUR_SLACK_CHANNEL_ID",
		slack.MsgOptionText(fmt.Sprintf("The item %s is expiring in %d days.", itemName, daysBeforeExpiration), false),
	)
	if err != nil {
		log.Println("Failed to send Slack notification:", err)
	} else {
		log.Println("Slack notification sent. Channel ID:", channelID)
	}
}
