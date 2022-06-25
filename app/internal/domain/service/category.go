package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go-grpc-template/internal/domain/models"
	"go-grpc-template/internal/repository"
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
func (c CategoryService) Create(ctx context.Context, request CreateCategoryInput) (*models.Category, error) {
	//валидация и создание сущности для записи в бд
	categoryFromRequest, err := c.NewCategoryFromRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	// запрос на создание
	if err := c.repo.Insert(ctx, *categoryFromRequest); err != nil {
		log.Error().Err(err).Msg("insert order")

		return nil, err
	}
	return categoryFromRequest, nil
}

// распарсиваем данные и приводим в к сущности Категория
func (c CategoryService) NewCategoryFromRequest(ctx context.Context, input CreateCategoryInput) (*models.Category, error) {

	// проверки, подгрузка значений из других моделей

	return &models.Category{
		//ID:                            uuid.New(),
		Name: input.Name,
	}, nil
}

// GetByID Получение запроса на платеж.
func (c CategoryService) GetByID(ctx context.Context, categoryID uuid.UUID) (*models.Category, error) {
	return c.repo.Get(ctx, &categoryID)
}

// GetAll Получение списка запросов на платеж.
func (c CategoryService) GetAll(ctx context.Context) ([]*models.Category, error) {

	return c.repo.List(ctx)
}
