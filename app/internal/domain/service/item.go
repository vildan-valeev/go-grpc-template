package service

import (
	"context"
	"github.com/google/uuid"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/internal/repository"
)

// item usecase - бизнес логика
type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

// Create Создание запроса на платеж.
func (c ItemService) Create(ctx context.Context, item models.Item) error {
	return c.repo.Insert(ctx, item)
}

// Get Получение запроса на платеж.
func (c ItemService) GetByID(ctx context.Context, itemID uuid.UUID) (*models.Item, error) {
	return c.repo.Get(ctx, &itemID)
}

// GetAll Получение списка запросов на платеж.
func (c ItemService) GetAll(ctx context.Context) ([]*models.Item, error) {

	return c.repo.List(ctx)
}
