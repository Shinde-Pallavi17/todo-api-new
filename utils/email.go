package utils

import (
	"fmt"
	"net/smtp"

	config "todo-manager/Config"
)

func SendReminderEmail(toEmail, title, body string) error {
	smtpConfig := config.LoadSMTPConfig()

	msg := "From: " + smtpConfig.From + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: Reminder: " + title + "\n\n" +
		body

	auth := smtp.PlainAuth(
		"",
		smtpConfig.User,
		smtpConfig.Pass,
		smtpConfig.Host,
	)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpConfig.Host, smtpConfig.Port),
		auth,
		smtpConfig.User,
		[]string{toEmail},
		[]byte(msg),
	)
}
