package service

import (
	"context"
	"github.com/google/uuid"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/internal/repository"
)

/*

Работа в бизнес логикой

*/

type Category interface {
	Create(ctx context.Context, category models.Category) error
	GetByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error)
	GetAll(ctx context.Context) ([]*models.Category, error)
}

type Item interface {
	Create(ctx context.Context, item models.Item) error
	GetByID(ctx context.Context, itemID uuid.UUID) (*models.Item, error)
	GetAll(ctx context.Context) ([]*models.Item, error)
}

type Services struct {
	Category Category
	Item     Item
}

type Deps struct {
	Repos *repository.Repositories
	Host  string
}

func NewServices(deps Deps) *Services {
	categoryService := NewCategoryService(deps.Repos.Category)
	itemService := NewItemService(deps.Repos.Item)

	return &Services{
		Category: categoryService,
		Item:     itemService,
	}
}
