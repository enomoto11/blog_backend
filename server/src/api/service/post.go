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
	CreatePost(ctx context.Context, rb request.POSTPostRequestBody) (*model.PostModel, error)
	FindAllPosts(ctx context.Context) ([]*model.PostModel, error)
}

type postService struct {
	postRepo     repository.PostRepository
	userRepo     repository.UserRepository
	categoryRepo repository.CategoryRepository
}

func NewPostService(
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
	categoryRepo repository.CategoryRepository,
) PostService {
	return &postService{
		postRepo,
		userRepo,
		categoryRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, rb request.POSTPostRequestBody) (*model.PostModel, error) {
	user, err := s.userRepo.FindByID(ctx, rb.UserID)
	if user == nil || err != nil {
		internalError := error2.NewInternalError(http.StatusBadRequest, err)
		return nil, internalError
	}

	category, err := s.categoryRepo.FindByID(ctx, rb.CategoryID)
	if category == nil || err != nil {
		internalError := error2.NewInternalError(http.StatusBadRequest, err)
		return nil, internalError
	}

	post, err := model.NewPost(
		model.NewPostTitle(rb.Title),
		model.NewPostBody(rb.Body),
		model.NewPostCategoryID(rb.CategoryID),
		model.NewPostUserID(rb.UserID),
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

func (s *postService) FindAllPosts(ctx context.Context) ([]*model.PostModel, error) {
	result, err := s.postRepo.FindAll(ctx)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}
