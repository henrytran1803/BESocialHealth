package ws

import (
	remindermodels "BESocialHealth/Internal/reminder_management/models"
	"BESocialHealth/component/appctx"
	"fmt"
	"strconv"
	"time"
)

type ReminderChecker struct {
	appCtx        appctx.AppContext
	wsManager     *WebSocketManager
	checkInterval time.Duration
}

func NewReminderChecker(appCtx appctx.AppContext, wsManager *WebSocketManager, checkInterval time.Duration) *ReminderChecker {
	return &ReminderChecker{
		appCtx:        appCtx,
		wsManager:     wsManager,
		checkInterval: checkInterval,
	}
}

func (rc *ReminderChecker) Start() {
	ticker := time.NewTicker(rc.checkInterval)
	go func() {
		for range ticker.C {
			rc.checkReminders()
		}
	}()
}

func (rc *ReminderChecker) checkReminders() {
	db := rc.appCtx.GetMainDBConnection()
	var reminders []remindermodels.Reminder
	now := time.Now()

	// Lấy các reminders cần thông báo
	if err := db.Where("date <= ? AND status = ?", now, 0).Find(&reminders).Error; err != nil {
		fmt.Println("Error fetching reminders:", err)
		return
	}
	fmt.Println(now)
	for _, reminder := range reminders {
		message := fmt.Sprintf("You have a reminder: %s", reminder.Description)
		rc.wsManager.SendToUser(fmt.Sprintf("%d", reminder.UserID), message)

		// Cập nhật trạng thái reminder
		reminder.Status = strconv.Itoa(1)
		db.Save(&reminder)
	}
}
