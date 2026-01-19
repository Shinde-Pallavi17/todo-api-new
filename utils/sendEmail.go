package utils

import (
	"fmt"
	"net/smtp"
	"strings"

	config "todo-manager/Config"
	"todo-manager/models"
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

func SendTaskAssignedEmail(
	toEmail string,
	toName string,
	task models.Task,
	assignerName string,
	assignerRole string,
) error {

	smtpConfig := config.LoadSMTPConfig()

	subject := "New Task Assigned"

	dueDate := "Not specified"
	if !task.DueDate.IsZero() {
		dueDate = task.DueDate.Format("02 Jan 2006")
	}

	body := fmt.Sprintf(
		`Hello %s,

A new task has been assigned to you.

Task: %s
Description: %s
Assigned by: %s (%s)
Due date: %s
Priority: %s
`,
		toName,
		task.Title,
		task.Description,
		assignerName,
		strings.Title(assignerRole),
		dueDate,
		strings.Title(task.Priority),
	)

	msg := "From: " + smtpConfig.From + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: " + subject + "\n\n" +
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
