package usermodels

type User struct {
	Id        int     `json:"id" gorm:"column:id"`
	Email     string  `gorm:"type:varchar(255);not null; column:email" json:"email"`
	FirstName string  `gorm:"type:varchar(255); column:firstname" json:"firstname"`
	LastName  string  `gorm:"type:varchar(255); column:lastname" json:"lastname"`
	Role      int     `json:"role"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	BDF       float64 `json:"bdf"`
	TDEE      float64 `json:"tdee"`
	Calorie   float64 `json:"calorie"`
	Status    int     `json:"status"`
}
type UserPhoto struct {
	Id        int     `json:"id" gorm:"column:id"`
	Email     string  `gorm:"type:varchar(255);not null; column:email" json:"email"`
	FirstName string  `gorm:"type:varchar(255); column:firstname" json:"firstname"`
	LastName  string  `gorm:"type:varchar(255); column:lastname" json:"lastname"`
	Role      int     `json:"role"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	BDF       float64 `json:"bdf"`
	TDEE      float64 `json:"tdee"`
	Calorie   float64 `json:"calorie"`
	Status    int     `json:"status"`
	Photo     *Photo  `json:"photo"`
}

type Photo struct {
	Id         int    `json:"id" gorm:"column:id"`
	Photo_type string `json:"photo_type" gorm:"column:photo_type"`
	Image      []byte `json:"image" gorm:"column:image"`
	Url        string `json:"url" gorm:"column:url"`
	User_id    string `json:"user_id" gorm:"column:user_id"`
}

func (User) TableName() string { return "users" }

type UserDetail struct {
	Id        int     `gorm:"primaryKey; column:id" json:"id"`
	Email     string  `gorm:"type:varchar(255);not null; column:email" json:"email"`
	FirstName string  `gorm:"type:varchar(255); column:firstname" json:"firstname"`
	LastName  string  `gorm:"type:varchar(255); column:lastname" json:"lastname"`
	Role      int     `json:"role"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	BDF       float64 `json:"bdf"`
	TDEE      float64 `json:"tdee"`
	Calorie   float64 `json:"calorie"`
	Status    int     `json:"status"`
	CreatedAt int64   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64   `json:"updated_at" gorm:"column:updated_at"`
}
