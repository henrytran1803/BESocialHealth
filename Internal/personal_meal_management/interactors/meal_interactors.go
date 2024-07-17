package mealinteractors

import (
	mealmodels "BESocialHealth/Internal/personal_meal_management/models"
	mealrepositories "BESocialHealth/Internal/personal_meal_management/repositories"
)

type MealInteractor struct {
	MealRepository *mealrepositories.MealRepository
}

func NewMealInteractor(repo *mealrepositories.MealRepository) *MealInteractor {
	return &MealInteractor{
		MealRepository: repo,
	}
}

func (i *MealInteractor) CreateMeal(meal *mealmodels.Meal) (*int, error) {
	id, err := i.MealRepository.CreateMeal(meal)
	if err != nil {
		return nil, err
	}
	return id, nil
}
func (i *MealInteractor) GetMeal(id string) (*mealmodels.GetMeal, error) {
	meal, err := i.MealRepository.GetMealById(id)
	if err != nil {
		return nil, err
	}
	return meal, nil
}
func (i *MealInteractor) GetMealByUserID(UserID string) (*[]mealmodels.GetMeal, error) {
	meal, err := i.MealRepository.GetMealByUserId(UserID)
	if err != nil {
		return nil, err
	}
	return meal, nil
}
func (i *MealInteractor) UpdateMealDetail(id int, meal *mealmodels.CreateMealDetail) error {
	if err := i.MealRepository.UpdateMealDetail(id, meal); err != nil {
		return err
	}
	return nil
}
func (i *MealInteractor) CreateMealDetail(mealDetail *mealmodels.CreateMealDetail) error {
	if err := i.MealRepository.CreateMealDetail(mealDetail); err != nil {
		return err
	}
	return nil
}
func (i *MealInteractor) DeleteMealById(id string) error {
	if err := i.MealRepository.DeleteMealById(id); err != nil {
		return err
	}
	return nil
}
func (i *MealInteractor) DeleteMealDetail(id string) error {
	if err := i.MealRepository.DeleteMealDetail(id); err != nil {
		return err
	}
	return nil
}
func (i *MealInteractor) GetMealByDate(id string, date string) (*mealmodels.GetMeal, error) {
	meal, err := i.MealRepository.GetMealByDate(&id, &date)
	if err != nil {
		return nil, err
	}
	if meal == nil {
		return nil, nil
	}
	return meal, nil
}
func (i *MealInteractor) GetInfomationCalories(id string, date string) (*float64, *float64, *float64, *mealmodels.MealNutrientTotals, error) {

	meal, schedule, calorie, nutritions, err := i.MealRepository.GetInformationCalories(id, date)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return meal, schedule, calorie, nutritions, nil
}
