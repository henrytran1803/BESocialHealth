package schedulerepositories

import (
	schedulemodels "BESocialHealth/Internal/personal_schedule_management/models"
	"errors"
	"gorm.io/gorm"
)

func (r *ScheduleRepository) GetListSchedule() ([]schedulemodels.ScheduleGet, error) {
	var results []schedulemodels.ScheduleGet
	var schedules []schedulemodels.Schedule
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Find(&schedules).Error; err != nil {
		return nil, err
	}
	for _, schedule := range schedules {
		var details []schedulemodels.ScheduleDetail
		if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", schedule.Id).Find(&details).Error; err != nil {
			return nil, err
		}
		result := schedulemodels.ScheduleGet{
			User_id: schedule.User_id,
			Time:    schedule.Time,
			Id:      schedule.Id,
			Detail:  details,
		}
		results = append(results, result)
	}
	return results, nil
}

func (r *ScheduleRepository) GetSchedule(scheduleID string) (*schedulemodels.ScheduleGet, error) {
	var schedule schedulemodels.Schedule
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("id = ?", scheduleID).Find(&schedule).Error; err != nil {
		return nil, err
	}
	var details []schedulemodels.ScheduleDetail
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", schedule.Id).Find(&details).Error; err != nil {
		return nil, err
	}
	result := schedulemodels.ScheduleGet{
		User_id: schedule.User_id,
		Time:    schedule.Time,
		Id:      schedule.Id,
		Detail:  details,
	}
	return &result, nil
}

func (r *ScheduleRepository) GetScheduleByUserId(userID string) (*[]schedulemodels.ScheduleGet, error) {
	var results []schedulemodels.ScheduleGet
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("user_id = ?", userID).Find(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *ScheduleRepository) GetScheduleDetailByScheduleId(scheduleID string) (*schedulemodels.ScheduleDetail, error) {
	var result schedulemodels.ScheduleDetail
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("id = ?", scheduleID).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ScheduleRepository) CreateSchedule(schedule *schedulemodels.ScheduleCreateFullOnly) (*int, error) {
	var scheduleDetails []schedulemodels.ScheduleDetailCreate
	var sche schedulemodels.ScheduleCreate

	// Gán giá trị cho các trường trong `sche`
	sche.Time = schedule.Time
	sche.User_id = schedule.User_id

	// Tạo một bản ghi mới trong bảng `schedules`
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Create(&sche).Error; err != nil {
		return nil, err
	}

	// Chuẩn bị dữ liệu cho bảng `schedule_details`
	for _, Detail := range schedule.Detail {
		scheduleDetail := schedulemodels.ScheduleDetailCreate{
			Schedule_id: sche.ID,
			Exersice_id: Detail.Exersice_id,
			Rep:         Detail.Rep,
			Time:        Detail.Time,
		}
		scheduleDetails = append(scheduleDetails, scheduleDetail)
	}

	// Tạo các bản ghi mới trong bảng `schedule_details`
	for _, scheduleDetail := range scheduleDetails {
		if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Create(&scheduleDetail).Error; err != nil {
			return nil, err
		}
	}

	return &sche.ID, nil
}

func (r *ScheduleRepository) CreateScheduleDetail(schedule *schedulemodels.ScheduleDetailCreate) error {
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Create(&schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *ScheduleRepository) UpdateSchedule(schedule *schedulemodels.ScheduleCreate) error {
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("id =?", schedule.ID).Save(&schedule).Error; err != nil {
		return err
	}
	return nil
}
func (r *ScheduleRepository) UpdateScheduleDetail(schedule *schedulemodels.ScheduleDetail) error {
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("id = ?", schedule.Id).Updates(&schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *ScheduleRepository) DeleteDetailById(id string) error {
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("id =?", id).Delete(&schedulemodels.ScheduleDetail{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *ScheduleRepository) DeleteDetailByScheduleId(scheduleId string) error {
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", scheduleId).Delete(&schedulemodels.ScheduleDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ScheduleRepository) DeleteScheduleById(id string) error {
	if err := r.DeleteDetailByScheduleId(id); err != nil {
		return err
	}

	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("id = ?", id).Delete(&schedulemodels.Schedule{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *ScheduleRepository) GetScheduleByDate(id *string, date *string) (*schedulemodels.ScheduleGet, error) {
	var schedule schedulemodels.ScheduleGet

	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("DATE(time) = ? AND user_id = ?", date, id).First(&schedule).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var schedule_details []schedulemodels.ScheduleDetail
	if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", schedule.Id).Find(&schedule_details).Error; err != nil {
		return nil, err
	}

	schedule.Detail = schedule_details
	return &schedule, nil
}
func (r *ScheduleRepository) GetScheduleDateToDate(fromdate, date *string, id *string) ([]schedulemodels.ScheduleGet, error) {
	var results []schedulemodels.ScheduleGet
	var schedules []schedulemodels.Schedule
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).
		Where("DATE(created_at) BETWEEN ? AND ? AND user_id = ?", fromdate, date, id).
		Find(&schedules).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	for _, schedule := range schedules {
		var details []schedulemodels.ScheduleDetail
		if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", schedule.Id).Find(&details).Error; err != nil {
			return nil, err
		}
		result := schedulemodels.ScheduleGet{
			User_id: schedule.User_id,
			Time:    schedule.Time,
			Id:      schedule.Id,
			Detail:  details,
		}
		results = append(results, result)
	}

	return results, nil

}
func (r *ScheduleRepository) GetListScheduleByUserId(userId string) ([]schedulemodels.ScheduleGet, error) {
	var results []schedulemodels.ScheduleGet
	var schedules []schedulemodels.Schedule
	if err := r.DB.Table(schedulemodels.Schedule{}.TableName()).Where("user_id = ?", userId).Find(&schedules).Error; err != nil {
		return nil, err
	}
	for _, schedule := range schedules {
		var details []schedulemodels.ScheduleDetail
		if err := r.DB.Table(schedulemodels.ScheduleDetail{}.TableName()).Where("schedule_id = ?", schedule.Id).Find(&details).Error; err != nil {
			return nil, err
		}
		result := schedulemodels.ScheduleGet{
			User_id:  schedule.User_id,
			Time:     schedule.Time,
			Id:       schedule.Id,
			Calories: schedule.Calories,
			Detail:   details,
		}
		results = append(results, result)
	}
	return results, nil
}
