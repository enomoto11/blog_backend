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
	CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.User, error)
	FindAllUser(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.User, error) {
	user, err := model.NewUser(
		model.NewUserFirstName(rb.First_name),
		model.NewUserLastName(rb.Last_name),
		model.NewUserEmail(rb.Email),
		model.NewUserPassword(rb.Password),
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

func (s *userService) FindAllUser(ctx context.Context) ([]*model.User, error) {
	users, errs := s.userRepo.FindAll(ctx)
	if errs != nil {
		internalError := error2.NewInternalError(http.StatusInternalServerError, errs[0])
		return nil, internalError
	}

	return users, nil
}
