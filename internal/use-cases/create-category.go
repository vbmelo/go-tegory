package use_cases

import (
	"github.com/vbmelo/go-tegory/internal/entities"
	"github.com/vbmelo/go-tegory/internal/repositories"
)

// lowercase sets the visibility to the current package (kinda like private in Java)
type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
