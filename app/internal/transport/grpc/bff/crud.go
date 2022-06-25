package bff

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"go-grpc-template/internal/domain/models"
	"go-grpc-template/internal/domain/service"
	errs "go-grpc-template/pkg/errors"
	"go-grpc-template/pkg/verify"
	pb "go-grpc-template/proto/generated"
	"time"
)

type CRUD struct {
	pb.UnimplementedCRUDServer

	services *service.Services

	Now func() time.Time
}

func NewCRUD(services *service.Services) *CRUD {
	return &CRUD{
		Now:      time.Now,
		services: services,
	}
}

// Create Создание Категории.
func (s *CRUD) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	createdRequest, err := GetValidatedCreateCategroyRequest(in)
	if err != nil {
		return nil, GRPCError(err)
	}

	createdCategory, err := s.services.Category.Create(ctx, *createdRequest)
	if err != nil {
		log.Error().Err(err).Msg("insert order")

		return nil, GRPCError(err)
	}

	return &pb.CreateCategoryResponse{
		CategoryId: createdCategory.ID.String(),
	}, nil
}

// Get Получение Категории.
func (s *CRUD) ReadCategory(ctx context.Context, in *pb.ReadCategoryRequest) (*pb.Category, error) {

	categoryID, err := uuid.Parse(in.GetCategoryId())
	if err != nil || categoryID == uuid.Nil {
		return nil, GRPCError(&errs.Error{Code: errs.InvalidID})
	}

	receivedCategory, err := s.services.Category.GetByID(ctx, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, GRPCError(&errs.Error{Code: errs.NotFound, Message: err.Error()})
		}

		return nil, GRPCError(err)
	}

	return receivedCategory.Proto()
}

// List Получение списка Категории.
func (s *CRUD) ListCategories(ctx context.Context, in *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	categories, err := s.services.Category.GetAll(ctx)
	if err != nil {
		return nil, GRPCError(err)
	}

	categoryList := models.Categories(categories)

	return &pb.ListCategoryResponse{Categories: categoryList.Proto()}, nil
}

func GetValidatedCreateCategroyRequest(in *pb.CreateCategoryRequest) (*service.CreateCategoryInput, error) {
	// grpc валидация запроса

	if err := verify.TextCheck(in.GetCategory().GetName()); err != nil && in.GetCategory().GetName() != "" {
		return nil, &errs.Error{Code: errs.InvalidName}
	}

	return getCreateCategoryRequest(in), nil
}
func getCreateCategoryRequest(in *pb.CreateCategoryRequest) *service.CreateCategoryInput {
	return &service.CreateCategoryInput{
		Name: in.GetCategory().GetName(),
	}
}

// CreateItem Создание запроса на платеж.
func (s *CRUD) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	createdRequest, err := GetValidatedCreateItemRequest(in)
	if err != nil {
		return nil, GRPCError(err)
	}
	createdItem, err := s.services.Item.Create(ctx, *createdRequest)
	if err != nil {
		log.Error().Err(err).Msg("insert order")

		return nil, GRPCError(err)
	}

	return &pb.CreateItemResponse{
		ItemId: createdItem.ID.String(),
	}, nil
}

// ReadItem Получение одного запроса на платеж.
func (s *CRUD) ReadItem(ctx context.Context, in *pb.ReadItemRequest) (*pb.Item, error) {

	itemID, err := uuid.Parse(in.GetItemId())
	if err != nil || itemID == uuid.Nil {
		return nil, GRPCError(&errs.Error{Code: errs.InvalidID})
	}

	receivedItem, err := s.services.Item.GetByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, GRPCError(&errs.Error{Code: errs.NotFound, Message: err.Error()})
		}
		return nil, GRPCError(err)
	}
	return receivedItem.Proto()
}

// ListItem Получение списка запросов на платеж.
func (s *CRUD) ListItem(ctx context.Context, in *pb.ListItemRequest) (*pb.ListItemResponse, error) {
	orders, err := s.services.Item.GetAll(ctx)
	if err != nil {
		return nil, GRPCError(err)
	}

	orderList := models.Items(orders)

	return &pb.ListItemResponse{Items: orderList.Proto()}, nil
}

func GetValidatedCreateItemRequest(in *pb.CreateItemRequest) (*service.CreateItemInput, error) {
	// grpc валидация запроса

	if err := verify.TextCheck(in.GetItem().GetName()); err != nil && in.GetItem().GetName() != "" {
		return nil, &errs.Error{Code: errs.InvalidName}
	}

	return getCreateItemRequest(in), nil
}
func getCreateItemRequest(in *pb.CreateItemRequest) *service.CreateItemInput {
	return &service.CreateItemInput{
		Name: in.GetItem().GetName(),
	}
}
