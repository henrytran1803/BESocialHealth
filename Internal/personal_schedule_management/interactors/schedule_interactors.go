package scheduleinteractors

import (
	schedulemodels "BESocialHealth/Internal/personal_schedule_management/models"
	schedulerepositories "BESocialHealth/Internal/personal_schedule_management/repositories"
)

type ScheduleInteractor struct {
	ScheduleRepository *schedulerepositories.ScheduleRepository
}

func NewScheduleInteractor(repo *schedulerepositories.ScheduleRepository) *ScheduleInteractor {
	return &ScheduleInteractor{
		ScheduleRepository: repo,
	}
}

func (i *ScheduleInteractor) CreateSchedule(schedule *schedulemodels.ScheduleCreateFull) error {
	if err := i.ScheduleRepository.CreateSchedule(schedule); err != nil {
		return err
	}
	return nil
}
func (i *ScheduleInteractor) CreateScheduleDetail(schedule *schedulemodels.ScheduleDetailCreate) error {
	if err := i.ScheduleRepository.CreateScheduleDetail(schedule); err != nil {
		return err
	}
	return nil
}
func (i *ScheduleInteractor) GetAllSchedule() ([]schedulemodels.ScheduleGet, error) {
	schedules, err := i.ScheduleRepository.GetListSchedule()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
func (i *ScheduleInteractor) GetScheduleById(id string) (*schedulemodels.ScheduleGet, error) {
	schedule, err := i.ScheduleRepository.GetSchedule(id)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
func (i *ScheduleInteractor) GetScheduleByUserId(userId string) (*[]schedulemodels.ScheduleGet, error) {
	schedules, err := i.ScheduleRepository.GetScheduleByUserId(userId)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
func (i *ScheduleInteractor) DeleteScheduleDetailById(id string) error {
	if err := i.ScheduleRepository.DeleteDetailById(id); err != nil {
		return err
	}
	return nil
}
func (i *ScheduleInteractor) DeleteScheduleById(id string) error {
	if err := i.ScheduleRepository.DeleteScheduleById(id); err != nil {
		return err
	}
	return nil
}

func (i *ScheduleInteractor) UpdateSchedule(schedule *schedulemodels.ScheduleCreate) error {
	if err := i.ScheduleRepository.UpdateSchedule(schedule); err != nil {
		return err
	}
	return nil
}
func (i *ScheduleInteractor) UpdateScheduleDetail(schedule *schedulemodels.ScheduleDetailCreate) error {
	if err := i.ScheduleRepository.UpdateScheduleDetail(schedule); err != nil {
		return err
	}
	return nil
}
