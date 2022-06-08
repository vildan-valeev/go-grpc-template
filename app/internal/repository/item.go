package repository

import (
	"context"
	"fmt"
	"go-bolvanka/internal/domain"
	"go-bolvanka/pkg/logging"
	"go-bolvanka/pkg/postgres"
)

type ItemRepository struct {
	pg     *postgres.Postgres
	logger *logging.Logger
}

func NewItemRepository(pg *postgres.Postgres, logger *logging.Logger) *ItemRepository {
	return &ItemRepository{pg: pg, logger: logger}
}

// GetAll Получение всех записей в базе данных
func (r *ItemRepository) GetAll(ctx context.Context) ([]domain.Item, error) {
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
	entities := make([]domain.Item, 0, _defaultEntityCap)

	for rows.Next() {
		e := domain.Item{}

		err = rows.Scan(&e.ID, &e.Name)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Post Создание записи в базе данных
//func (r ItemRepository) Post(ctx context.Context, item *item.Item) error {
//	im.ID = user.ID
//
//	model := toModel(bm)
//
//	res, err := r.db.InsertOne(ctx, model)
//	if err != nil {
//		return err
//	}
//
//	bm.ID = res.InsertedID.(primitive.ObjectID).Hex()
//	return nil
//}

//// GetOne Получение всех записей в базе данных
//func (r ItemRepository) GetOne(ctx context.Context, id string) (item.Item, error) {
//	q := `
//		SELECT id, name, age FROM public.book;
//	`
//
//	rows, err := r.client.Query(ctx, q)
//	if err != nil {
//		return nil, err
//	}
//
//	books := make([]book.Book, 0)
//
//	for rows.Next() {
//		var bk Book
//
//		err = rows.Scan(&bk.ID, &bk.Name, &bk.Age)
//		if err != nil {
//			return nil, err
//		}
//
//		books = append(books, bk.ToDomain())
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return books, nil
//}
//
//// Update Получение всех записей в базе данных
//func (r ItemRepository) Update(ctx context.Context, category item.Item) error {
//	q := `
//		SELECT id, name, age FROM public.book;
//	`
//
//	rows, err := r.client.Query(ctx, q)
//	if err != nil {
//		return nil, err
//	}
//
//	books := make([]book.Book, 0)
//
//	for rows.Next() {
//		var bk Book
//
//		err = rows.Scan(&bk.ID, &bk.Name, &bk.Age)
//		if err != nil {
//			return nil, err
//		}
//
//		books = append(books, bk.ToDomain())
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return books, nil
//}
//
//// Delete Удаление строки из базы данных
//func (r ItemRepository) Delete(ctx context.Context, id string) error {
//	objID, _ := primitive.ObjectIDFromHex(id)
//	uID, _ := primitive.ObjectIDFromHex(user.ID)
//
//	_, err := r.db.DeleteOne(ctx, bson.M{"_id": objID, "userId": uID})
//	return err
//}
