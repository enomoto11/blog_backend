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
	CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.POSTUserModel, error)
	FindAllUsers(ctx context.Context) ([]*model.POSTUserModel, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.POSTUserModel, error) {
	user, err := model.NewPOSTUser(
		model.NewPOSTUserFirstName(rb.First_name),
		model.NewPOSTUserLastName(rb.Last_name),
		model.NewPOSTUserEmail(rb.Email),
		model.NewPOSTUserPassword(rb.Password),
	)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	result, err := s.userRepo.Create(ctx, user)
	if err != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}

func (s *userService) FindAllUsers(ctx context.Context) ([]*model.POSTUserModel, error) {
	users, errs := s.userRepo.FindAll(ctx)
	if errs != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, errs[0])
		return nil, internalError
	}

	return users, nil
}
