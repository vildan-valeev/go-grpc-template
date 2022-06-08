package service

import (
	"context"
	"go-bolvanka/internal/domain"
)

type Category interface {
	//CreateCategory(ctx context.Context, category *Category) error
	AllCategories(ctx context.Context) ([]domain.Category, error)
	//DeleteCategory(ctx context.Context, id string) error
}

// ItemUseCase - методы для работы с бизнес -логикой
type Item interface {
	//CreateItem(ctx context.Context, item *Item) error
	AllItems(ctx context.Context) ([]domain.Item, error)
	//DeleteItem(ctx context.Context, id string) error
}
