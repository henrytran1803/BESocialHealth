package mealrepositories

import (
	mealmodels "BESocialHealth/Internal/personal_meal_management/models"
	"gorm.io/gorm"
)

func (r *MealRepository) CreateMealOnly(meal *mealmodels.CreateMeal) error {
	return r.DB.Table(mealmodels.CreateMeal{}.TableName()).Create(&meal).Error
}
func (r *MealRepository) CreateMealDetail(meal *mealmodels.CreateMealDetail) error {
	return r.DB.Table(mealmodels.CreateMealDetail{}.TableName()).Create(&meal).Error
}

func (r *MealRepository) CreateMeal(meals *mealmodels.Meal) (*int, error) {
	meal := mealmodels.CreateMeal{
		UserId: meals.UserId,
	}
	if err := r.DB.Create(&meal).Error; err != nil {
		return nil, err
	}
	mealID := meal.ID
	for _, dish := range meals.Dishes {
		mealDetail := mealmodels.CreateMealDetail{
			MealId:  mealID,
			DishId:  dish.Id,
			Serving: dish.Serving,
		}
		if err := r.DB.Create(&mealDetail).Error; err != nil {
			return nil, err
		}
	}
	return &mealID, nil
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

func (r *MealRepository) GetMealByDate(id *string, date *string) (*mealmodels.GetMeal, error) {
	var meal mealmodels.GetMeal
	if err := r.DB.Table(mealmodels.GetMeal{}.TableName()).Where("DATE(date) = ? AND user_id = ?", date, id).First(&meal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var dishes []mealmodels.GetDish
	if err := r.DB.Table(mealmodels.GetDish{}.TableName()).Where("meal_id = ?", meal.ID).Find(&dishes).Error; err != nil {
		return nil, err
	}

	meal.Dishes = dishes
	return &meal, nil
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

func (r *MealRepository) GetInformationCalories(userID string, date string) (*float64, *float64, *float64, *mealmodels.MealNutrientTotals, error) {
	var mealCalorie float64
	var scheduleCalorie float64
	var userCalorie float64
	var nutrientTotals mealmodels.MealNutrientTotals

	if err := r.DB.Table("meals").
		Where("user_id = ?", userID).
		Where("DATE(date) = DATE(?)", date).
		Select("COALESCE(SUM(total_calorie), 0)").
		Scan(&mealCalorie).Error; err != nil {
		return nil, nil, nil, nil, err
	}

	if err := r.DB.Table("schedules").
		Where("user_id = ?", userID).
		Where("DATE(time) = DATE(?)", date).
		Select("COALESCE(SUM(calories_burn), 0)").
		Scan(&scheduleCalorie).Error; err != nil {
		return nil, nil, nil, nil, err
	}

	if err := r.DB.Table("users").
		Where("id = ?", userID).
		Select("calorie").
		Scan(&userCalorie).Error; err != nil {
		return nil, nil, nil, nil, err
	}

	if err := r.DB.Table("meal_detail").
		Joins("JOIN dishes ON meal_detail.dish_id = dishes.id").
		Where("meal_detail.meal_id IN (SELECT id FROM meals WHERE user_id = ? AND DATE(date) = DATE(?))", userID, date).
		Select(`
        COALESCE(SUM((dishes.calorie * meal_detail.serving) / dishes.serving), 0) AS TotalCalorie,
        COALESCE(SUM(dishes.protein * meal_detail.serving / dishes.serving), 0) AS TotalProtein,
        COALESCE(SUM(dishes.fat * meal_detail.serving / dishes.serving), 0) AS TotalFat,
        COALESCE(SUM(dishes.carb * meal_detail.serving / dishes.serving), 0) AS TotalCarb,
        COALESCE(SUM(dishes.sugar * meal_detail.serving / dishes.serving), 0) AS TotalSugar`).
		Scan(&nutrientTotals).Error; err != nil {
		return nil, nil, nil, nil, err
	}
	remainingCalorie := userCalorie - mealCalorie + scheduleCalorie

	return &mealCalorie, &scheduleCalorie, &remainingCalorie, &nutrientTotals, nil
}
