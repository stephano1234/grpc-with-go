package service

import (
	"context"
	"io"

	"github.com/stephano1234/grpc-go/internal/database"
	"github.com/stephano1234/grpc-go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryResquest) (*pb.CreateCategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, &in.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{
		Category: &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: *category.Description,
		},
	}, nil
}

func (c *CategoryService) QueryCategory(ctx context.Context, in *pb.Blank) (*pb.QueryCategoryResponse, error) {
	categories, err := c.CategoryDB.GetAll()
	if err != nil {
		return nil, err
	}
	var categoriesResponse []*pb.Category 
	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: *category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	return &pb.QueryCategoryResponse{
		Categories: categoriesResponse,
	}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.GetByID(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: *category.Description,
	}, nil
}

func (c *CategoryService) UpdateCategory(ctx context.Context, in *pb.Category) (*pb.Blank, error) {
	if err := c.CategoryDB.UpdateByID(in.Id, in.Name, in.Description); err != nil {
		return nil, err
	}
	return &pb.Blank{}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) (error) {
	var categoriesResponse []*pb.Category
	for {
		categoryResquest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.QueryCategoryResponse{
				Categories: categoriesResponse,
			})
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.Create(categoryResquest.Name, &categoryResquest.Description)
		if err != nil {
			return err
		}
		categoriesResponse = append(categoriesResponse, &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: *category.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) (error) {
	for {
		categoryResquest, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.Create(categoryResquest.Name, &categoryResquest.Description)
		if err != nil {
			return err
		}
		categoryResponse := &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: *category.Description,
		}
		stream.Send(categoryResponse)
	}
}
