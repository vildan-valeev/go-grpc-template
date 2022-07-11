package domain

import (
	"github.com/google/uuid"
	"go-grpc-template/pkg/utils"
	pb "go-grpc-template/proto/generated"
)

type Item struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}

// ответ за запрос ReadCategory из gRPC
func (c *Item) Proto() (*pb.Item, error) {

	return &pb.Item{
		Id:   utils.ReplaceNilUUID(c.ID),
		Name: c.Name,
	}, nil
}

// массив - список категорий
type Items []*Item

// ответ за запрос ListCategories из gRPC
func (is *Items) Proto() []*pb.Item {
	if len(*is) == 0 {
		return nil
	}

	items := make([]*pb.Item, 0, len(*is))
	for _, i := range *is {
		items = append(items, &pb.Item{
			Id:   i.ID.String(),
			Name: i.Name,
			//...
		})
	}

	return items
}
