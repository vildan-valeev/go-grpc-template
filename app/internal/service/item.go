package service

import (
	"context"
	"go-bolvanka/internal/domain"
	"go-bolvanka/internal/repository"
)

// item usecase - бизнес логика
type ItemService struct {
	itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
	}
}

//func (i ItemUseCase) CreateItem(ctx context.Context, item *item.Item) error {
//	return i.itemRepo.Post(ctx, item)
//}

func (i *ItemService) AllItems(ctx context.Context) ([]domain.Item, error) {
	return i.itemRepo.GetAll(ctx)
}

//func (i ItemUseCase) DeleteItem(ctx context.Context, id string) error {
//	return i.itemRepo.Delete(ctx, id)
//}
