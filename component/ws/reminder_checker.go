package ws

import (
	remindermodels "BESocialHealth/Internal/reminder_management/models"
	"BESocialHealth/component/appctx"
	"fmt"
	"time"
)

// ReminderChecker kiểm tra các reminders và gửi thông báo
type ReminderChecker struct {
	appCtx        appctx.AppContext
	wsManager     *WebSocketManager
	checkInterval time.Duration
}

// NewReminderChecker tạo một ReminderChecker mới
func NewReminderChecker(appCtx appctx.AppContext, wsManager *WebSocketManager, checkInterval time.Duration) *ReminderChecker {
	return &ReminderChecker{
		appCtx:        appCtx,
		wsManager:     wsManager,
		checkInterval: checkInterval,
	}
}

// Start bắt đầu kiểm tra reminders định kỳ
func (rc *ReminderChecker) Start() {
	ticker := time.NewTicker(rc.checkInterval)
	go func() {
		for range ticker.C {
			rc.checkReminders()
		}
	}()
}

// checkReminders kiểm tra các reminders và gửi thông báo
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
	// Gửi thông báo đến từng user
	for _, reminder := range reminders {
		message := fmt.Sprintf("You have a reminder: %s", reminder.Description)
		rc.wsManager.SendToUser(fmt.Sprintf("%d", reminder.UserID), message)

		// Cập nhật trạng thái reminder
		reminder.Status = 1
		db.Save(&reminder)
	}
}
