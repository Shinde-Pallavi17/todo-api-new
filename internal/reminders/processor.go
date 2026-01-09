package reminders

import (
	"log"
	"time"

	"todo-manager/models"

	"gorm.io/gorm"
)

func ProcessReminders(db *gorm.DB) {
	var reminders []models.Reminder
	now := time.Now().UTC()

	err := db.
		Where(`
			reminder_at <= ?
			AND status = ?
			AND triggered_at IS NULL
			AND retry_count < ?
		`, now, "pending", 3).
		Limit(20).
		Find(&reminders).Error

	if err != nil {
		log.Println("failed to fetch reminders:", err)
		return
	}

	for _, r := range reminders {
		if err := SendNotification(r); err != nil {
			db.Model(&r).Updates(map[string]interface{}{
				"retry_count": gorm.Expr("retry_count + 1"),
			})

			if r.RetryCount+1 >= 3 {
				db.Model(&r).Update("status", "failed")
			}
			continue
		}

		triggeredAt := time.Now().UTC()
		db.Model(&r).Updates(map[string]interface{}{
			"status":       "sent",
			"triggered_at": &triggeredAt,
		})
	}
}
