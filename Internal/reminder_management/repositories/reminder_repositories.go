package reminderrepositories

import remindermodels "BESocialHealth/Internal/reminder_management/models"

func (r *ReminderRepository) CreateReminder(reminder *remindermodels.ReminderCreate) error {
	return r.DB.Table(remindermodels.Reminder{}.TableName()).Create(reminder).Error
}

func (r *ReminderRepository) UpdateReminder(reminder *remindermodels.ReminderCreate) error {
	return r.DB.Table(remindermodels.Reminder{}.TableName()).Save(reminder).Error
}

func (r *ReminderRepository) DeleteReminder(id int) error {
	return r.DB.Table(remindermodels.Reminder{}.TableName()).Delete(&remindermodels.Reminder{}, id).Error
}

func (r *ReminderRepository) GetReminderByID(id int) (*remindermodels.Reminder, error) {
	var reminder remindermodels.Reminder
	err := r.DB.Table(remindermodels.Reminder{}.TableName()).First(&reminder, id).Error
	return &reminder, err
}
func (r *ReminderRepository) GetReminderByUserID(id int) (*[]remindermodels.Reminder, error) {
	var reminder []remindermodels.Reminder
	err := r.DB.Table(remindermodels.Reminder{}.TableName()).Where("user_id = ?", id).Find(&reminder).Error
	return &reminder, err
}
