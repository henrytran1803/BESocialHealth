package exersicemodels

import "BESocialHealth/comon"

//CREATE TABLE `exersice_type` (
//`id` int NOT NULL AUTO_INCREMENT,
//`name` varchar(255),
//`status` int,
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`deleted_at` TIMESTAMP NULL DEFAULT NULL,
//PRIMARY KEY (`id`)
//);
//
//CREATE TABLE `exersices` (
//`id` int NOT NULL AUTO_INCREMENT,
//`exersice_type` int NOT NULL,
//`name` varchar(255),
//`description` text,
//`calorie` double NOT NULL,
//`rep_serving` int,
//`time_serving` int,
//`status` int,
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`deleted_at` TIMESTAMP NULL DEFAULT NULL,
//PRIMARY KEY (`id`),
//FOREIGN KEY (`exersice_type`) REFERENCES `exersice_type`(`id`)
//);

type Exersice struct {
	comon.SQLModel
	Name          string        `json:"name" gorm:"type:varchar(255);not null; column:name" form:"name" binding:"required"`
	Description   string        `json:"description" gorm:"type:varchar(255);not null" form:"description" binding:"required"`
	Calorie       float64       `json:"calorie" gorm:"type:decimal(10,2);not null" form:"calorie"`
	Rep_serving   int           `json:"rep_serving" gorm:"type:int(11);not null" form:"rep_serving"`
	Time_serving  int           `json:"time_serving" gorm:"type:int(11);not null" form:"time_serving"`
	Exersice_type int           `json:"exersice_type"`
	ExersiceType  Exersice_type `gorm:"foreignKey:Exersice_type"`
}

type Exersice_type struct {
	comon.SQLModel
	Name string `json:"name" gorm:"type:varchar(255);not null" form:"name"`
}

func (Photo) TableName() string { return "photos" }

func (Exersice) TableName() string      { return "exersices" }
func (Exersice_type) TableName() string { return "exersice_type" }

type Photo struct {
	comon.SQLModel
	Photo_type  string `json:"photo_type" gorm:"column:photo_type"`
	Image       []byte `json:"image" gorm:"column:image"`
	Url         string `json:"url" gorm:"column:url"`
	Exersice_id string `json:"exersice_id" gorm:"column:exersice_id"`
}

type CreateExersice struct {
	Name          string  `json:"name" gorm:"type:varchar(255);not null; column:name" form:"name"`
	Description   string  `json:"description" gorm:"type:varchar(255);not null" form:"description"`
	Calorie       float64 `json:"calorie" gorm:"type:decimal(10,2);not null" form:"calorie"`
	Rep_serving   int     `json:"rep_serving" gorm:"type:int(11);not null" form:"rep_serving"`
	Time_serving  int     `json:"time_serving" gorm:"type:int(11);not null" form:"time_serving"`
	Exersice_type int     `json:"exersice_type" gorm:"type:int(11);not null" form:"exersice_type"`
	Image         []byte  `json:"image" gorm:"column:image" form:"image" `
}
type GetExersiceList struct {
	comon.SQLModel
	Name          string        `json:"name" gorm:"type:varchar(255);not null; column:name" form:"name" binding:"required"`
	Description   string        `json:"description" gorm:"type:varchar(255);not null;column:description" form:"description" binding:"required"`
	Calorie       float64       `json:"calorie" gorm:"type:decimal(10,2);not null;column:calorie" form:"calorie"`
	Rep_serving   int           `json:"rep_serving" gorm:"type:int(11);not null;column:rep_serving" form:"rep_serving"`
	Time_serving  int           `json:"time_serving" gorm:"type:int(11);not null;column:time_serving" form:"time_serving"`
	Exersice_type Exersice_type `json:"exersice_type"`
	Photo         []Photo       `json:"photo" gorm:"foreignKey:exersice_id;not null;" form:"photo" binding:"required"`
}
