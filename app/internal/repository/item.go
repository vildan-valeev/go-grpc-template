package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/pkg/database"
)

type ItemRepository struct {
	*database.DB
}

func NewItemRepository(db *database.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (s *ItemRepository) Insert(ctx context.Context, item models.Item) error {
	tx, err := s.Pool().Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := s.insertItem(ctx, item, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *ItemRepository) insertItem(ctx context.Context, item models.Item, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `INSERT INTO item (name) VALUES ($1)`, item.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *ItemRepository) List(ctx context.Context) ([]*models.Item, error) {

	sql := `SELECT id, name FROM item`

	rows, err := s.Pool().Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	itemsModel, err := s.fetch(ctx, rows)
	if err != nil {
		return nil, err
	}

	return itemsModel, nil
}

func (s *ItemRepository) fetch(ctx context.Context, rows pgx.Rows) ([]*models.Item, error) {
	items := make([]*models.Item, 0)

	for rows.Next() {
		var item models.Item

		if err := rows.Scan(
			&item.ID,
			&item.Name,
		); err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return items, nil
}

func (s *ItemRepository) Get(ctx context.Context, itemID *uuid.UUID) (*models.Item, error) {
	var item models.Item

	sql := `SELECT id, name FROM item WHERE id = $1`

	if err := pgxscan.Get(ctx, s.Pool(), &item, sql, itemID); err != nil {
		return nil, err
	}

	return &item, nil

}
