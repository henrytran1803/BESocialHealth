package interactors

import (
	exersicemodels "BESocialHealth/Internal/exersice_management/models"
	exersicerepositories "BESocialHealth/Internal/exersice_management/repositories"
	"fmt"
)

type ExersiceInteractor struct {
	ExersiceRepository *exersicerepositories.ExersiceRepository
}

func NewExersiceInteractor(repo *exersicerepositories.ExersiceRepository) *ExersiceInteractor {
	return &ExersiceInteractor{
		ExersiceRepository: repo,
	}
}

func (i *ExersiceInteractor) CreateExersice(exersice *exersicemodels.Exersice, imageData []byte, fileName string) error {
	err := i.ExersiceRepository.CreateExersice(exersice)
	if err != nil {
		return err
	}
	exersiceID := exersice.Id
	photo := &exersicemodels.Photo{
		Photo_type:  "1",
		Image:       imageData,
		Url:         fileName,
		Exersice_id: fmt.Sprintf("%d", exersiceID),
	}

	err = i.ExersiceRepository.CreatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}

func (i *ExersiceInteractor) UpdateExersice(id int, exersice *exersicemodels.Exersice, imageData []byte, fileName string) error {
	if exersice == nil {
		return fmt.Errorf("exersice must not be nil")
	}
	if err := i.ExersiceRepository.UpdateExersice(id, exersice); err != nil {
		return err
	}
	exersiceID := exersice.Id
	photo := &exersicemodels.Photo{
		Photo_type:  "1",
		Image:       imageData,
		Url:         fileName,
		Exersice_id: fmt.Sprintf("%d", exersiceID),
	}
	err := i.ExersiceRepository.UpdatePhoto(id, photo)
	if err != nil {
		return err
	}
	return nil
}
func (i *ExersiceInteractor) DeleteExersice(id int) error {
	if err := i.ExersiceRepository.DeleteExersiceById(id); err != nil {
		return err
	}
	if err := i.ExersiceRepository.DeletePhotoByExersice(id); err != nil {
		return err
	}
	return nil
}
func (i *ExersiceInteractor) GetExersice(id int) (*exersicemodels.Exersice, error) {
	exersice, err := i.ExersiceRepository.FindExersiceById(id)
	if err != nil {
		return nil, err
	}
	return &exersice, nil
}
func (i *ExersiceInteractor) GetAllExersice() ([]exersicemodels.GetExersiceList, error) {
	exersices, err := i.ExersiceRepository.GetListExersice()
	if err != nil {
		return nil, err
	}
	return exersices, nil
}
