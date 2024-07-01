package personalcontentrepositories

import "gorm.io/gorm"

type PersonalContentRepository struct {
	DB *gorm.DB
}

func NewPersonalContentRepository(db *gorm.DB) *PersonalContentRepository {
	return &PersonalContentRepository{DB: db}
}
