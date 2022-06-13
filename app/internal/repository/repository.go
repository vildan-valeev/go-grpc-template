package repository

import (
	"github.com/google/uuid"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/pkg/database"
)

import (
	"context"
)

/*
Работа с БД
*/

// Category - методы для работы с БД
type Category interface {
	Insert(ctx context.Context, c models.Category) error
	List(ctx context.Context) ([]*models.Category, error)
	Get(ctx context.Context, categoryID *uuid.UUID) (*models.Category, error)
}

// Item - методы для работы с БД
type Item interface {
	Insert(ctx context.Context, i models.Item) error
	List(ctx context.Context) ([]*models.Item, error)
	Get(ctx context.Context, itemID *uuid.UUID) (*models.Item, error)
}

type Repositories struct {
	Category Category
	Item     Item
}

// создаем структуру репозиториев
func NewRepositories(db *database.DB) *Repositories {
	return &Repositories{
		Category: NewCategoryRepository(db),
		Item:     NewItemRepository(db),
	}
}
