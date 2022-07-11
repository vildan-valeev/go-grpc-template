package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go-grpc-template/internal/domain"
	"go-grpc-template/internal/repository"
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
func (i ItemService) Create(ctx context.Context, request CreateItemInput) (*domain.Item, error) {
	//валидация и создание сущности для записи в бд
	itemFromRequest, err := i.NewItemFromRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	// запрос на создание
	if err := i.repo.Insert(ctx, *itemFromRequest); err != nil {
		log.Error().Err(err).Msg("insert order")

		return nil, err
	}
	return itemFromRequest, nil
}

// распарсиваем данные и приводим в к сущности Категория
func (i ItemService) NewItemFromRequest(ctx context.Context, input CreateItemInput) (*domain.Item, error) {

	// проверки, подгрузка значений из других моделей

	return &domain.Item{
		//ID:                            uuid.New(),
		Name: input.Name,
	}, nil
}

// Get Получение запроса на платеж.
func (i ItemService) GetByID(ctx context.Context, itemID uuid.UUID) (*domain.Item, error) {
	return i.repo.Get(ctx, &itemID)
}

// GetAll Получение списка запросов на платеж.
func (i ItemService) GetAll(ctx context.Context) ([]*domain.Item, error) {

	return i.repo.List(ctx)
}
