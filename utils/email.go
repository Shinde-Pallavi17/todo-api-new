package utils

import (
	"fmt"
	"net/smtp"
	config "todo-manager/Config"
)

func SendReminderEmail(toEmail, title, body string) error {
	msg := "From: " + config.FromEmail + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: Reminder: " + title + "\n\n" +
		body

	auth := smtp.PlainAuth(
		"",
		config.SMTPUser,
		config.SMTPPass,
		config.SMTPHost,
	)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort),
		auth,
		config.SMTPUser,
		[]string{toEmail},
		[]byte(msg),
	)
}
