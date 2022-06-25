package service

import (
	"context"
	"github.com/google/uuid"
	"go-grpc-template/internal/domain/models"
	"go-grpc-template/internal/repository"
)

/*

Работа в бизнес логикой

*/

type (
	Category interface {
		Create(ctx context.Context, request CreateCategoryInput) (*models.Category, error)
		GetByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error)
		GetAll(ctx context.Context) ([]*models.Category, error)
	}
	CreateCategoryInput struct {
		ID   uuid.UUID
		Name string
	}
)

type (
	Item interface {
		Create(ctx context.Context, request CreateItemInput) (*models.Item, error)
		GetByID(ctx context.Context, itemID uuid.UUID) (*models.Item, error)
		GetAll(ctx context.Context) ([]*models.Item, error)
	}
	CreateItemInput struct {
		ID   uuid.UUID
		Name string
	}
)

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
