package reminders

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func StartReminderWorker(db *gorm.DB) {
	log.Println("Reminder worker started")

	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			ProcessReminders(db)
		}
	}()
}
