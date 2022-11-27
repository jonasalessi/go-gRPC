package categoryService

import (
	"context"
	"errors"
	"io"
	mapper "my-grpc/m/application/mapper"
	"my-grpc/m/infra/database"
	"my-grpc/m/infra/pb"
)

var CategoryRequiredFieldsError = errors.New("Name and Description are required!")

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDB
}

func NewCategoryService(db database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (c *CategoryService) CreateCategory(stream pb.CategoryService_CreateCategoryServer) error {
	for {
		categoryRequest, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if categoryRequest.Name == "" || categoryRequest.Description == "" {
			return CategoryRequiredFieldsError
		}
		categoryDb := c.CategoryDB.Save(categoryRequest.Name, categoryRequest.Description)
		err = stream.Send(mapper.ToPbCategoryFrom(categoryDb))
		if err != nil {
			return err
		}
	}
}

func (c *CategoryService) GetCategoryBy(ctx context.Context, id *pb.Id) (*pb.Category, error) {
	categoryDb, err := c.CategoryDB.GetById(id.Id)
	if err != nil {
		return nil, err
	}
	return mapper.ToPbCategoryFrom(categoryDb), nil
}

func (c *CategoryService) GetAllCategories(ctx context.Context, b *pb.Blank) (*pb.CategoryList, error) {
	categories := c.CategoryDB.GetAll()
	output := make([]*pb.Category, 0)
	for _, categoryDb := range categories {
		output = append(output, mapper.ToPbCategoryFrom(categoryDb))
	}
	return &pb.CategoryList{
		Categories: output,
	}, nil
}
