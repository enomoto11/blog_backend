package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.User) (*model.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}

func (r *userRepository) Create(ctx context.Context, m *model.User) (*model.User, error) {
	entity, err := r.client.User.Create().
		SetID(m.GetID()).
		SetFirstName(m.GetFirstName()).
		SetLastName(m.GetLastName()).
		SetEmail(m.GetEmail()).
		SetPassword(m.GetPassword()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return teamModelFromEntity(entity)
}

func teamModelFromEntity(entity *ent.User) (*model.User, error) {
	opts := []model.NewUserOption{
		model.NewUserID(entity.ID),
		model.NewUserFirstName(entity.FirstName),
		model.NewUserLastName(entity.LastName),
		model.NewUserEmail(entity.Email),
		model.NewUserPassword(entity.Password),
	}

	return model.NewUser(opts...)
}
