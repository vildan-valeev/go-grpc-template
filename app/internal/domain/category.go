package domain

import (
	"github.com/google/uuid"
	"go-grpc-template/pkg/utils"

	pb "go-grpc-template/proto/generated"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// ответ за запрос ReadCategory из gRPC
func (c *Category) Proto() (*pb.Category, error) {

	return &pb.Category{
		Id:   utils.ReplaceNilUUID(c.ID),
		Name: c.Name,
	}, nil
}

// массив - список категорий
type Categories []*Category

// ответ за запрос ListCategories из gRPC
func (cs *Categories) Proto() []*pb.Category {
	if len(*cs) == 0 {
		return nil
	}

	categories := make([]*pb.Category, 0, len(*cs))
	for _, cat := range *cs {
		categories = append(categories, &pb.Category{
			Id:   cat.ID.String(),
			Name: cat.Name,
			//...
		})
	}

	return categories
}
