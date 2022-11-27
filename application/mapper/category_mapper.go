package mapperCategory

import (
	"my-grpc/m/infra/database"
	"my-grpc/m/infra/pb"
)

func ToPbCategoryFrom(category *database.CategoryEntity) *pb.Category {
	return &pb.Category{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}
}
