package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"go-grpc-template/internal/domain"
	"go-grpc-template/pkg/database"
)

type CategoryRepository struct {
	*database.DB
}

func NewCategoryRepository(db *database.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (s *CategoryRepository) Insert(ctx context.Context, category domain.Category) error {
	tx, err := s.Pool().Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := s.insertCategory(ctx, category, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *CategoryRepository) insertCategory(ctx context.Context, category domain.Category, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `INSERT INTO category (name) VALUES ($1)`, category.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryRepository) List(ctx context.Context) ([]*domain.Category, error) {

	sql := `SELECT id, name FROM category`

	rows, err := s.Pool().Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ordersModel, err := s.fetch(ctx, rows)
	if err != nil {
		return nil, err
	}

	return ordersModel, nil
}
func (s *CategoryRepository) fetch(ctx context.Context, rows pgx.Rows) ([]*domain.Category, error) {
	categories := make([]*domain.Category, 0)

	for rows.Next() {
		var category domain.Category

		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return categories, nil
}

func (s *CategoryRepository) Get(ctx context.Context, categoryID *uuid.UUID) (*domain.Category, error) {
	var category domain.Category

	sql := `SELECT id, name FROM category WHERE id = $1`

	if err := pgxscan.Get(ctx, s.Pool(), &category, sql, categoryID); err != nil {
		return nil, err
	}

	return &category, nil
}
