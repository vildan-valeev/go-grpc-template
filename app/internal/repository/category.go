package repository

import (
	"context"
	"fmt"
	"go-bolvanka/internal/domain"
	"go-bolvanka/pkg/logging"
	"go-bolvanka/pkg/postgres"
)

type CategoryRepository struct {
	pg     *postgres.Postgres
	logger *logging.Logger
}

func NewCategoryRepository(pg *postgres.Postgres, logger *logging.Logger) *CategoryRepository {
	return &CategoryRepository{pg: pg, logger: logger}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	// формируем sql запрос через Builder
	sql, _, err := r.pg.Builder.
		Select("id, name").
		From("category").
		ToSql()
	r.logger.Info(fmt.Sprintf("SQL Query: %s", sql))
	// проверяем составленный запрос
	if err != nil {
		return nil, fmt.Errorf("CategoryRepository - GetAll - r.pg.Builder: %w", err)
	}
	// заходим в Pool, делаем запрос
	rows, err := r.pg.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("CategoryRepository - GetAll - r.Pool.Query: %w", err)
	}
	// закрытие по выходу из функции
	defer rows.Close()
	// распарсиваем записи в сущности
	entities := make([]domain.Category, 0, _defaultEntityCap)

	for rows.Next() {
		e := domain.Category{}

		err = rows.Scan(&e.ID, &e.Name)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

//func (r *CategoryRepository) Post(ctx context.Context, category *category.Category) error {
//	q := `
//		INSERT INTO author
//		    (name, age)
//		VALUES
//		       ($1, $2)
//		RETURNING id
//	`
//	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
//	if err := r.client.QueryRow(ctx, q, category.Name, 123).Scan(&category.ID); err != nil {
//		var pgErr *pgconn.PgError
//		if errors.As(err, &pgErr) {
//			pgErr = err.(*pgconn.PgError)
//			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
//			r.logger.Error(newErr)
//			return newErr
//		}
//		return err
//	}
//
//	return nil
//}

//func (r *CategoryRepository) GetOne(ctx context.Context, id string) (category.Category, error) {
//	q := `
//		SELECT id, name FROM public.author WHERE id = $1
//	`
//	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
//
//	var ath category.Category
//	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
//	if err != nil {
//		return category.Category{}, err
//	}
//
//	return ath, nil
//}
//
//func (r *CategoryRepository) Update(ctx context.Context, category category.Category) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
//
//	panic("implement me")
//}
