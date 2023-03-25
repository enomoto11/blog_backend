package service

import (
	"blog/api/controller/request"
	error2 "blog/api/error"
	"blog/api/model"
	"blog/api/repository"
	"context"
	"net/http"
)

type UserService interface {
	CreateUser(ctx context.Context, rb request.POSTUserRequestBody) (*model.POSTUserModel, error)
	FindAllUsers(ctx context.Context) ([]*model.GETUserModel, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, rb request.POSTUserRequestBody) (*model.POSTUserModel, error) {
	user, err := model.NewPOSTUser(
		model.NewPOSTUserFirstName(rb.FirstName),
		model.NewPOSTUserLastName(rb.LastName),
		model.NewPOSTUserEmail(rb.Email),
		model.NewPOSTUserPassword(rb.Password),
	)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusBadRequest, err)
		return nil, internalError
	}

	result, err := s.userRepo.Create(ctx, user)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}

func (s *userService) FindAllUsers(ctx context.Context) ([]*model.GETUserModel, error) {
	users, err := s.userRepo.FindAll(ctx)

	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return users, nil
}
