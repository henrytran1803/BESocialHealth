package reminderinteractors

import (
	remindermodels "BESocialHealth/Internal/reminder_management/models"
	reminderrepositories "BESocialHealth/Internal/reminder_management/repositories"
)

type ReminderInteractor struct {
	ReminderRepository *reminderrepositories.ReminderRepository
}

func NewReminderInteractor(repo *reminderrepositories.ReminderRepository) *ReminderInteractor {
	return &ReminderInteractor{
		ReminderRepository: repo,
	}
}
func (i *ReminderInteractor) CreateReminder(reminder *remindermodels.ReminderCreate) error {
	if err := i.ReminderRepository.CreateReminder(reminder); err != nil {
		return err
	}
	return nil
}

func (i *ReminderInteractor) UpdateReminder(reminder *remindermodels.ReminderCreate) error {
	if err := i.ReminderRepository.UpdateReminder(reminder); err != nil {
		return err
	}
	return nil
}
func (i *ReminderInteractor) DeleteReminderById(id int) error {
	if err := i.ReminderRepository.DeleteReminder(id); err != nil {
		return err
	}
	return nil
}
func (i *ReminderInteractor) GetReminderById(id int) (*remindermodels.Reminder, error) {
	remind, err := i.ReminderRepository.GetReminderByID(id)
	if err != nil {
		return nil, err
	}
	return remind, nil
}
func (i *ReminderInteractor) GetAllReminderById(id int) (*[]remindermodels.Reminder, error) {
	remind, err := i.ReminderRepository.GetReminderByUserID(id)
	if err != nil {
		return nil, err
	}
	return remind, nil
}
