package service

import (
	"blog/api/controller/request"
	error2 "blog/api/error"
	"blog/api/model"
	"blog/api/repository"
	"context"
	"net/http"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, rb request.POSTCategoryRequestBody) (*model.CategoryModel, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, rb request.POSTCategoryRequestBody) (*model.CategoryModel, error) {
	category, err := model.NewCategoryBeforeCreated(
		model.NewCategoryName(rb.Name),
	)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusBadRequest, err)
		return nil, internalError
	}

	result, err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}
