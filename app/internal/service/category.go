package service

import (
	"context"
	"go-bolvanka/internal/domain"
	"go-bolvanka/internal/repository"
)

type CategoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

//func (c CategoryUseCase) CreateCategory(ctx context.Context, category *category.Category) error {
//
//	return c.categoryRepo.Post(ctx, category)
//}

func (c CategoryService) AllCategories(ctx context.Context) ([]domain.Category, error) {
	return c.categoryRepo.GetAll(ctx)
}

//func (c CategoryUseCase) DeleteCategory(ctx context.Context, id string) error {
//	return c.categoryRepo.Delete(ctx, id)
//}
