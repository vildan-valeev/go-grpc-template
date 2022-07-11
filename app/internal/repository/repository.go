package repository

import (
	"context"
	"github.com/google/uuid"
	"go-grpc-template/internal/domain"
	"go-grpc-template/pkg/database"
)

/*
Работа с БД
*/

// Category - методы для работы с БД
type Category interface {
	Insert(ctx context.Context, c domain.Category) error
	List(ctx context.Context) ([]*domain.Category, error)
	Get(ctx context.Context, categoryID *uuid.UUID) (*domain.Category, error)
}

// Item - методы для работы с БД
type Item interface {
	Insert(ctx context.Context, i domain.Item) error
	List(ctx context.Context) ([]*domain.Item, error)
	Get(ctx context.Context, itemID *uuid.UUID) (*domain.Item, error)
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
