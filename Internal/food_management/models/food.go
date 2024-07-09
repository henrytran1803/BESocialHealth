package foodmodels

// CREATE TABLE `dishes` (
// `id` int NOT NULL AUTO_INCREMENT,
// `name` varchar(255) NOT NULL,
// `description` text,
// `calorie` double,
// `protein` double,
// `fat` double,
// `carb` double,
// `sugar` double,
// `serving` double,
// `status` int,
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// `deleted_at` TIMESTAMP NULL DEFAULT NULL,
// PRIMARY KEY (`id`)
// );

//CREATE TABLE `photo_type` (
//`id` int NOT NULL AUTO_INCREMENT,
//`name` varchar(255),
//`status` int,
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`deleted_at` TIMESTAMP NULL DEFAULT NULL,
//PRIMARY KEY (`id`)
//);
//CREATE TABLE `photos` (
//`id` int NOT NULL AUTO_INCREMENT,
//`photo_type` int,
//`url` varchar(255),
//`post_id` int,
//`comment_id` int,
//`exersice_id` int,
//`dish_id` int,
//`user_id` int,
//`status` int,
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`deleted_at` TIMESTAMP NULL DEFAULT NULL,
//PRIMARY KEY (`id`),
//FOREIGN KEY (`photo_type`) REFERENCES `photo_type`(`id`),
//FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
//FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`),
//FOREIGN KEY (`comment_id`) REFERENCES `comments`(`id`),
//FOREIGN KEY (`dish_id`) REFERENCES `dishes`(`id`),
//FOREIGN KEY (`exersice_id`) REFERENCES `exersices`(`id`)
//);

type Food struct {
	Id          int     `json:"id" gorm:"column:id"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Calorie     float64 `json:"calorie" gorm:"column:calorie"`
	Protein     float64 `json:"protein" gorm:"column:protein"`
	Fat         float64 `json:"fat" gorm:"column:fat"`
	Carb        float64 `json:"carb" gorm:"column:carb"`
	Sugar       float64 `json:"sugar" gorm:"column:sugar"`
	Serving     int     `json:"serving" gorm:"column:serving"`
}
type GetFood struct {
	Id          int     `json:"id" gorm:"column:id"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Calorie     float64 `json:"calorie" gorm:"column:calorie"`
	Protein     float64 `json:"protein" gorm:"column:protein"`
	Fat         float64 `json:"fat" gorm:"column:fat"`
	Carb        float64 `json:"carb" gorm:"column:carb"`
	Sugar       float64 `json:"sugar" gorm:"column:sugar"`
	Serving     int     `json:"serving" gorm:"column:serving"`
	Photos      []Photo `json:"photos" gorm:"foreignKey:Dish_id"`
}

func (Food) TableName() string  { return "dishes" }
func (Photo) TableName() string { return "photos" }

type Photo struct {
	Id         int    `json:"id" gorm:"column:id"`
	Photo_type string `json:"photo_type" gorm:"column:photo_type"`
	Image      []byte `json:"image" gorm:"column:image"`
	Url        string `json:"url" gorm:"column:url"`
	Dish_id    string `json:"dish_id" gorm:"column:dish_id"`
}
type FoodCreate struct {
	Name        string  `form:"name" json:"name" gorm:"column:name"`
	Description string  `form:"description" json:"description" gorm:"column:description"`
	Calorie     float64 `form:"calorie" json:"calorie" gorm:"column:calorie"`
	Protein     float64 `form:"protein" json:"protein" gorm:"column:protein"`
	Fat         float64 `form:"fat" json:"fat" gorm:"column:fat"`
	Carb        float64 `form:"carb" json:"carb" gorm:"column:carb"`
	Sugar       float64 `form:"sugar" json:"sugar" gorm:"column:sugar"`
	Serving     int     `form:"serving" json:"serving" gorm:"column:serving"`
	Image       []byte  `form:"photos" json:"photos" `
}
type FoodUpdate struct {
	Id          int     `json:"id" gorm:"column:id"`
	Name        string  `form:"name" json:"name" gorm:"column:name"`
	Description string  `form:"description" json:"description" gorm:"column:description"`
	Calorie     float64 `form:"calorie" json:"calorie" gorm:"column:calorie"`
	Protein     float64 `form:"protein" json:"protein" gorm:"column:protein"`
	Fat         float64 `form:"fat" json:"fat" gorm:"column:fat"`
	Carb        float64 `form:"carb" json:"carb" gorm:"column:carb"`
	Sugar       float64 `form:"sugar" json:"sugar" gorm:"column:sugar"`
	Serving     int     `form:"serving" json:"serving" gorm:"column:serving"`
}
type PhotoBase struct {
	Id          int     `json:"id" gorm:"column:id"`
	Photo_type  string  `json:"photo_type" gorm:"column:photo_type"`
	Image       []byte  `json:"image" gorm:"column:image"`
	Url         string  `json:"url" gorm:"column:url"`
	Dish_id     *string `json:"dish_id" gorm:"column:dish_id"`
	Exersice_id *string `json:"exersice_id" gorm:"column:exersice_id"`
	Post_id     *string `json:"post_id" gorm:"column:post_id"`
	Coment_id   *string `json:"comment_id" gorm:"column:comment_id"`
	User_id     *string `json:"user_id" gorm:"column:user_id"`
}
