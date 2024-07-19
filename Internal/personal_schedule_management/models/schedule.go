package schedulemodels

import (
	"time"
)

// CREATE TABLE `schedules` (
// `id` int NOT NULL AUTO_INCREMENT,
// `user_id` int,
// `time` timestamp,
// `calories_burn` double,
// `status` int,
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// `deleted_at` TIMESTAMP NULL DEFAULT NULL,
// PRIMARY KEY (`id`),
// FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
// );
//
// CREATE TABLE `schedule_detail` (
// `id` int NOT NULL AUTO_INCREMENT,
// `schedule_id` int NOT NULL,
// `exersice_id` int NOT NULL,
// `rep` int,
// `time` int,
// `status` int,
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// `deleted_at` TIMESTAMP NULL DEFAULT NULL,
// PRIMARY KEY (`id`),
// UNIQUE KEY `unique_schedule` (`schedule_id`, `exersice_id`),
// FOREIGN KEY (`schedule_id`) REFERENCES `schedules`(`id`),
// FOREIGN KEY (`exersice_id`) REFERENCES `exersices`(`id`)
// );
type Schedule struct {
	Id       int        `json:"id" gorm:"column:id"`
	User_id  int64      `json:"user_id" gorm:"column:user_id;"`
	Time     *time.Time `json:"time" gorm:"column:time;"`
	Calories float32    `json:"calories" gorm:"column:calories;"`
}
type ScheduleDetail struct {
	Id          int `json:"id" gorm:"column:id"`
	Schedule_id int `json:"schedule_id" gorm:"column:schedule_id;"`
	Exersice_id int `json:"exersice_id" gorm:"column:exersice_id;"`
	Rep         int `json:"rep" gorm:"column:rep"`
	Time        int `json:"time" gorm:"column:time"`
}

func (Schedule) TableName() string       { return "schedules" }
func (ScheduleDetail) TableName() string { return "schedule_detail" }

type ScheduleCreateFull struct {
	User_id int64                  `json:"user_id" gorm:"column:user_id;not null"`
	Time    time.Time              `json:"time" gorm:"column:time"`
	Detail  []ScheduleDetailCreate `json:"detail"`
}

type ScheduleCreate struct {
	ID      int        `json:"id" gorm:"column:id;"`
	User_id int64      `json:"user_id" gorm:"column:user_id;not null"`
	Time    *time.Time `json:"time" gorm:"column:time"`
}
type ScheduleDetailCreate struct {
	//ID          int `json:"id" gorm:"column:id;"`
	Schedule_id int `json:"schedule_id" gorm:"column:schedule_id;not null"`
	Exersice_id int `json:"exersice_id" gorm:"column:exersice_id;not null"`
	Rep         int `json:"rep" gorm:"column:rep"`
	Time        int `json:"time" gorm:"column:time"`
}
type ScheduleCreateFullOnly struct {
	User_id int64                      `json:"user_id" gorm:"column:user_id;not null"`
	Time    *time.Time                 `json:"time" gorm:"column:time"`
	Detail  []ScheduleDetailCreateOnly `json:"detail"`
}
type ScheduleDetailCreateOnly struct {
	Exersice_id int `json:"exersice_id" gorm:"column:exersice_id;not null"`
	Rep         int `json:"rep" gorm:"column:rep"`
	Time        int `json:"time" gorm:"column:time"`
}

type ScheduleGet struct {
	Id        int              `json:"id" gorm:"column:id"`
	User_id   int64            `json:"user_id" gorm:"column:user_id;"`
	Time      *time.Time       `json:"time" gorm:"column:time;not null"`
	Calories  float32          `json:"calories" gorm:"column:calories_burn"`
	Status    int              `json:"status" gorm:"column:status"`
	Create_at *time.Time       `json:"create_at" gorm:"column:created_at"`
	Detail    []ScheduleDetail `json:"detail" gorm:"foreignKey:schedule_id;references:id"`
}
