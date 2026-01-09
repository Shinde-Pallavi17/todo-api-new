package reminders

import (
	"errors"
	"fmt"
	"log"
	config "todo-manager/Config"
	"todo-manager/models"
	"todo-manager/utils"
)

func SendNotification(r models.Reminder) error {
	var user models.User
	var task models.Task

	//Fetch user
	if err := config.DB.First(&user, r.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Email == "" {
		return errors.New("user email missing")
	}

	if !user.EmailVerified {
		return errors.New("email not verified")
	}

	//Fetch task
	if err := config.DB.First(&task, r.TaskID).Error; err != nil {
		return errors.New("task not found")
	}

	//Send email
	subject := task.Title
	body := fmt.Sprintf(
		"Reminder for your task:\n\nTitle: %s\nDescription: %s\nDue Date: %s",
		task.Title,
		task.Description,
		task.DueDate.Format("2006-01-02"),
	)

	if err := utils.SendReminderEmail(user.Email, subject, body); err != nil {
		return err
	}

	//SUCCESS LOG
	log.Printf(
		"Reminder email sent | user=%s task_id=%d reminder_id=%d",
		user.Email,
		task.Id,
		r.ID,
	)

	return nil
}
