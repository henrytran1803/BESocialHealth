package exersicerepositories

import (
	exersicemodels "BESocialHealth/Internal/exersice_management/models"
	foodmodels "BESocialHealth/Internal/food_management/models"
)

func (r *ExersiceRepository) GetListExersice() ([]exersicemodels.GetExersiceList, error) {

	var exersices []exersicemodels.Exersice
	var getExersices []exersicemodels.GetExersiceList
	if err := r.DB.Find(&exersices).Error; err != nil {
		return nil, err
	}
	for _, exersice := range exersices {
		var photos []exersicemodels.Photo
		if err := r.DB.Where("exersice_id = ?", exersice.Id).Find(&photos).Error; err != nil {
			return nil, err
		}
		getExersice := exersicemodels.GetExersiceList{
			Name:         exersice.Name,
			Description:  exersice.Description,
			Calorie:      exersice.Calorie,
			Rep_serving:  exersice.Rep_serving,
			Time_serving: exersice.Time_serving,
			SQLModel:     exersice.SQLModel,
			Photo:        photos,
		}
		getExersices = append(getExersices, getExersice)
	}
	return getExersices, nil
}

func (r *ExersiceRepository) CreateExersice(exersice *exersicemodels.Exersice) error {
	return r.DB.Table(exersicemodels.Exersice{}.TableName()).Create(&exersice).Error
}
func (r *ExersiceRepository) CheckExistExersiceByName(name string) bool {
	var exersice exersicemodels.Exersice
	if err := r.DB.Where("name = ?", name).First(&exersice).Error; err != nil {
		return false
	}
	return true
}

func (r *ExersiceRepository) FindExersiceById(id int) (exersicemodels.Exersice, error) {
	var exersice exersicemodels.Exersice
	if err := r.DB.Where("id = ?", id).First(&exersice).Error; err != nil {
		return exersice, err
	}
	return exersice, nil
}
func (r *ExersiceRepository) UpdateExersice(id int, exersice *exersicemodels.Exersice) error {
	exersice.Id = id
	return r.DB.Table(exersicemodels.Exersice{}.TableName()).Updates(exersice).Error
}

func (r *ExersiceRepository) DeleteExersiceById(id int) error {
	var exersice exersicemodels.Exersice
	if err := r.DB.Where("id = ?", id).Delete(&exersice).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExersiceRepository) CreatePhoto(photo *exersicemodels.Photo) error {
	photo.Photo_type = "1"
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Create(photo).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExersiceRepository) UpdatePhoto(id int, photo *exersicemodels.Photo) error {
	// Use Updates to update the record with the specific exersice_id
	return r.DB.Table(exersicemodels.Photo{}.TableName()).Where("exersice_id = ?", id).Updates(photo).Error
}

func (r *ExersiceRepository) DeletePhotoById(id int) error {
	if err := r.DB.Table(exersicemodels.Photo{}.TableName()).Where("id = ?", id).Delete(&exersicemodels.Photo{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExersiceRepository) DeletePhotoByExersice(exerciseId int) error {
	if err := r.DB.Table(exersicemodels.Photo{}.TableName()).Where("exercise_id = ?", exerciseId).Delete(&exersicemodels.Photo{}).Error; err != nil {
		return err
	}
	return nil
}
