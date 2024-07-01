package mealrepositories

import mealmodels "BESocialHealth/Internal/personal_meal_management/models"

func (r *MealRepository) CreateMealOnly(meal *mealmodels.CreateMeal) error {
	return r.DB.Table(mealmodels.CreateMeal{}.TableName()).Create(&meal).Error
}
func (r *MealRepository) CreateMealDetail(meal *mealmodels.CreateMealDetail) error {
	return r.DB.Table(mealmodels.CreateMealDetail{}.TableName()).Create(&meal).Error
}

func (r *MealRepository) CreateMeal(meals *mealmodels.Meal) error {
	meal := mealmodels.CreateMeal{
		UserId: meals.UserId,
	}
	if err := r.DB.Create(&meal).Error; err != nil {
		return err
	}
	mealID := meal.ID
	for _, dish := range meals.Dishes {
		mealDetail := mealmodels.CreateMealDetail{
			MealId:  mealID,
			DishId:  dish.Id,
			Serving: dish.Serving,
		}
		if err := r.DB.Create(&mealDetail).Error; err != nil {
			return err
		}
	}
	return nil
}
func (r *MealRepository) GetMealById(id string) (*mealmodels.GetMeal, error) {
	var meal mealmodels.GetMeal
	if err := r.DB.Preload("Dishes").First(&meal, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &meal, nil
}
func (r *MealRepository) GetMealByUserId(userId string) (*[]mealmodels.GetMeal, error) {
	var meals []mealmodels.GetMeal
	if err := r.DB.Table(mealmodels.GetMeal{}.TableName()).Where("user_id = ?", userId).Find(&meals).Error; err != nil {
		return nil, err
	}
	for i := range meals {
		var dishes []mealmodels.GetDish
		if err := r.DB.Table(mealmodels.GetDish{}.TableName()).Where("meal_id = ?", meals[i].ID).Find(&dishes).Error; err != nil {
			return nil, err
		}
		meals[i].Dishes = dishes
	}

	return &meals, nil
}
func (r *MealRepository) UpdateMealDetail(mealDetailId int, mealDetail *mealmodels.CreateMealDetail) error {
	return r.DB.Table(mealmodels.CreateMealDetail{}.TableName()).Where("id = ?", mealDetailId).Updates(&mealDetail).Error
}
func (r *MealRepository) DeleteMealDetail(mealDetailId string) error {
	return r.DB.Table(mealmodels.CreateMealDetail{}.TableName()).Where("id = ?", mealDetailId).Delete(&mealmodels.CreateMealDetail{}).Error
}
func (r *MealRepository) DeleteMealById(id string) error {
	if err := r.DB.Table(mealmodels.CreateMealDetail{}.TableName()).Where("meal_id = ?", id).Delete(&mealmodels.CreateMealDetail{}).Error; err != nil {
		return err
	}
	if err := r.DB.Table(mealmodels.CreateMeal{}.TableName()).Where("id = ?", id).Delete(&mealmodels.CreateMeal{}).Error; err != nil {

		return err

	}
	return nil
}
