package repositories

import "github.com/vbmelo/go-tegory/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
