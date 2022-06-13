package service

import (
	"context"
	"github.com/google/uuid"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/internal/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// Create Создание запроса на платеж.
func (c CategoryService) Create(ctx context.Context, category models.Category) error {
	return c.repo.Insert(ctx, category)
}

// GetByID Получение запроса на платеж.
func (c CategoryService) GetByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
	return c.repo.Get(ctx, &categoryID)
}

// GetAll Получение списка запросов на платеж.
func (c CategoryService) GetAll(ctx context.Context) ([]*models.Category, error) {

	return c.repo.List(ctx)
}
