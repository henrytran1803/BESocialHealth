package foodinteractors

import (
	foodmodels "BESocialHealth/Internal/food_management/models"
	foodrepositories "BESocialHealth/Internal/food_management/repositories"
	"fmt"
	"strconv"
)

type FoodInteractor struct {
	FoodRepository *foodrepositories.FoodRepository
}

func NewFoodInteractor(repo *foodrepositories.FoodRepository) *FoodInteractor {
	return &FoodInteractor{
		FoodRepository: repo,
	}
}

func (i *FoodInteractor) GetFood(id string) (*foodmodels.Food, error) {
	food, err := i.FoodRepository.GetFoodById(id)
	if err != nil {
		return nil, err
	}
	return food, nil
}
func (i *FoodInteractor) DeleteFood(id int) error {
	err := i.FoodRepository.DeleteFoodById(strconv.Itoa(id))
	if err != nil {
		return err
	}

	if err := i.FoodRepository.DeletePhotoByFood(strconv.Itoa(id)); err != nil {
		return err
	}
	return nil
}
func (i *FoodInteractor) CreateFood(food *foodmodels.Food, imageData []byte, fileName string) error {
	err := i.FoodRepository.CreateFood(food)
	if err != nil {
		return err
	}
	foodID := food.Id
	photo := &foodmodels.Photo{
		Photo_type: "1",
		Image:      imageData,
		Url:        fileName,
		Dish_id:    fmt.Sprintf("%d", foodID),
	}

	err = i.FoodRepository.CreatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}
func (i *FoodInteractor) UpdateFood(food *foodmodels.Food, imageData []byte, fileName string) error {
	err := i.FoodRepository.UpdateFoodById(food)
	if err != nil {
		return err
	}
	foodID := food.Id
	if len(imageData) > 0 && fileName != "" {
		photo := &foodmodels.Photo{
			Photo_type: "1",
			Image:      imageData,
			Url:        fileName,
			Dish_id:    fmt.Sprintf("%d", foodID),
		}
		err = i.FoodRepository.UpdatePhoto(photo)
		if err != nil {
			return err
		}
	}
	return nil
}
func (i *FoodInteractor) UpdateFoodNonePhoto(food *foodmodels.FoodUpdate) error {
	err := i.FoodRepository.UpdateFood(food)
	if err != nil {
		return err
	}
	return nil
}
func (i *FoodInteractor) GetListFood() ([]foodmodels.GetFood, error) {
	foods, err := i.FoodRepository.GetListFood()
	if err != nil {
		return nil, err
	}
	return foods, nil
}
func (i *FoodInteractor) DeletePhotoById(id int) error {
	err := i.FoodRepository.DeletePhotoById(strconv.Itoa(id))
	if err != nil {
		return err
	}
	if err := i.FoodRepository.DeletePhotoByFood(strconv.Itoa(id)); err != nil {
		return err
	}
	return nil
}
func (i *FoodInteractor) CreatePhoto(photo *foodmodels.PhotoBase) error {
	err := i.FoodRepository.CreatePhotoBase(photo)
	if err != nil {
		return err
	}
	return nil
}
func (i *FoodInteractor) CreateListPhoto(photos []foodmodels.PhotoBase) error {
	err := i.FoodRepository.CreatePhotoListBase(photos)
	if err != nil {
		return err
	}
	return nil
}
