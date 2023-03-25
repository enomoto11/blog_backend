package service

import (
	"blog/api/controller/request"
	error2 "blog/api/error"
	"blog/api/model"
	"blog/api/repository"
	"context"
	"net/http"
)

type PostService interface {
	CreatePost(ctx context.Context, rb request.POSTPostRequestBody) (*model.POSTPostModel, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, rb request.POSTPostRequestBody) (*model.POSTPostModel, error) {
	post, err := model.NewPOSTPost(
		model.NewPOSTPostTitle(rb.Title),
		model.NewPOSTPostBody(rb.Body),
		model.NewPOSTPostCategoryID(rb.CategoryID),
		model.NewPOSTPostUserID(rb.UserID),
	)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusBadRequest, err)
		return nil, internalError
	}

	result, err := s.postRepo.Create(ctx, post)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}
