package foodrepositories

import (
	foodmodels "BESocialHealth/Internal/food_management/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *FoodRepository) CreateFood(food *foodmodels.Food) error {
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Create(food).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) GetFoodById(id string) (*foodmodels.Food, error) {
	var food foodmodels.Food
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Where("id = ?", id).Find(&food).Error; err != nil {
		return nil, err
	}
	return &food, nil
}
func (r *FoodRepository) CheckExistFoodByName(name string) (bool, error) {
	var food foodmodels.Food
	err := r.DB.Table(food.TableName()).Where("name = ?", name).First(&food).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *FoodRepository) GetAllFood() ([]foodmodels.Food, error) {
	var foods []foodmodels.Food
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
func (r *FoodRepository) DeleteFoodById(id string) error {
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Where("id = ?", id).Delete(&foodmodels.Food{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) GetListFood() ([]foodmodels.GetFood, error) {
	var foods []foodmodels.Food
	var getFoods []foodmodels.GetFood
	if err := r.DB.Find(&foods).Error; err != nil {
		return nil, err
	}

	for _, food := range foods {
		var photos []foodmodels.Photo
		if err := r.DB.Where("dish_id = ?", food.Id).Find(&photos).Error; err != nil {
			return nil, err
		}
		getFood := foodmodels.GetFood{
			Id:          food.Id,
			Name:        food.Name,
			Description: food.Description,
			Calorie:     food.Calorie,
			Protein:     food.Protein,
			Fat:         food.Fat,
			Carb:        food.Carb,
			Sugar:       food.Sugar,
			Serving:     food.Serving,
			Photos:      photos,
		}
		getFoods = append(getFoods, getFood)
	}

	return getFoods, nil
}

func (r *FoodRepository) UpdateFoodById(food *foodmodels.Food) error {
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Save(food).Error; err != nil {
		return err
	}
	return nil
}

//Photo

func (r *FoodRepository) CreatePhoto(photo *foodmodels.Photo) error {
	photo.Photo_type = "1"
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Create(photo).Error; err != nil {
		return err
	}
	return nil
}

func (r *FoodRepository) UpdatePhoto(photo *foodmodels.Photo) error {
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Save(photo).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) DeletePhotoByFood(id string) error {
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Delete(&foodmodels.Photo{}, "dish_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) DeletePhotoById(id string) error {
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Delete(&foodmodels.Photo{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) CreatePhotoBase(photo *foodmodels.PhotoBase) error {
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Create(photo).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) CreatePhotoListBase(photos []foodmodels.PhotoBase) error {
	for _, photo := range photos {
		if err := r.CreatePhotoBase(&photo); err != nil {
			return err
		}
	}
	return nil
}
func (r *FoodRepository) UpdatePhotoBase(photo *foodmodels.PhotoBase) error {
	if err := r.DB.Table(foodmodels.Photo{}.TableName()).Save(photo).Error; err != nil {
		return err
	}
	return nil
}
func (r *FoodRepository) UpdateFood(food *foodmodels.FoodUpdate) error {
	var count int64
	err := r.DB.Table("meal_detail").Where("dish_id = ?", food.Id).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("cannot update food because it is referenced in mealdetail")
	}
	if err := r.DB.Table(foodmodels.Food{}.TableName()).Save(food).Error; err != nil {
		return err
	}
	return nil
}
