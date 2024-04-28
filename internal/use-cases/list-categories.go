package use_cases

import (
	"github.com/vbmelo/go-tegory/internal/entities"
	"github.com/vbmelo/go-tegory/internal/repositories"
)

// lowercase sets the visibility to the current package (kinda like private in Java)
type listCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoryUseCase(repository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
