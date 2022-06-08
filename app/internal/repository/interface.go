package repository

import (
	"context"
	"go-bolvanka/internal/domain"
)

const _defaultEntityCap = 64

// Item - методы для работы с БД
type Item interface {
	//Post(ctx context.Context, item *Item) error
	GetAll(ctx context.Context) ([]*domain.Item, error)
	//GetOne(ctx context.Context, id string) (Item, error)
	//Update(ctx context.Context, item Item) error
	//Delete(ctx context.Context, id string) error
}

// Category - методы для работы с БД
type Category interface {
	//Post(ctx context.Context, category *Category) error
	GetAll(ctx context.Context) ([]domain.Category, error)
	//GetOne(ctx context.Context, id string) (Category, error)
	//Update(ctx context.Context, category Category) error
	//Delete(ctx context.Context, id string) error
}
