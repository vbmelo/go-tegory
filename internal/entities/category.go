package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Business Rules

	return category, nil
}

func (c *Category) IsValid() error {

	if len(c.Name) < 5 {
		return fmt.Errorf("name must be at least 5 characters, got %d", len(c.Name))
	}

	return nil
}
