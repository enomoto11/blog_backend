package service

import (
	"blog/api/controller/request"
	"blog/api/error"
	"blog/api/model"
	"blog/api/repository"
	"context"
	"net/http"
)

type UserService interface {
	CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.User, *error.InternalError)
}

type userService struct {
	createUserRepo repository.UserRepository
}

func NewUserService(createUserRepo repository.UserRepository) UserService {
	return &userService{
		createUserRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, rb request.CreateUserRequestBody) (*model.User, *error.InternalError) {
	team, err := model.NewUser(
		model.NewUserFirstName(rb.FirstName),
		model.NewUserLastName(rb.LastName),
		model.NewUserEmail(rb.Email),
		model.NewUserPassword(rb.Password),
	)
	if err != nil {
		internalError := error.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	result, err := s.createUserRepo.Create(ctx, team)
	if err != nil {
		internalError := error.NewInternalError(http.StatusInternalServerError, err)
		return nil, internalError
	}

	return result, nil
}
