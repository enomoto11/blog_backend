package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.POSTUserModel) (*model.POSTUserModel, error)
	FindAll(ctx context.Context) ([]*model.GETUserModel, []error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}

func (r *userRepository) Create(ctx context.Context, m *model.POSTUserModel) (*model.POSTUserModel, error) {
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

func (r *userRepository) FindAll(ctx context.Context) ([]*model.GETUserModel, []error) {
	entities, err := r.client.User.Query().All(ctx)

	if err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}

	return teamModelsFromEntities(entities)
}

func teamModelFromEntity(entity *ent.User) (*model.POSTUserModel, error) {
	opts := []model.NewPOSTUserOption{
		model.NewPOSTUserID(entity.ID),
		model.NewPOSTUserFirstName(entity.FirstName),
		model.NewPOSTUserLastName(entity.LastName),
		model.NewPOSTUserEmail(entity.Email),
		model.NewPOSTUserPassword(entity.Password),
	}

	return model.NewPOSTUser(opts...)
}

func teamModelsFromEntities(entities []*ent.User) ([]*model.GETUserModel, []error) {
	var results []*model.GETUserModel
	var errs []error

	for _, entity := range entities {
		opts := []model.NewGETUserOption{
			model.NewGETUserID(entity.ID),
			model.NewGETUserFirstName(entity.FirstName),
			model.NewGETUserLastName(entity.LastName),
			model.NewGETUserEmail(entity.Email),
		}
		result, err := model.NewGETUser(opts...)
		results = append(results, result)
		errs = append(errs, err)
	}

	return results, errs
}
