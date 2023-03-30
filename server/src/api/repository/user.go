package repository

import (
	"blog/api/model"
	"blog/ent"
	"blog/ent/user"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, m *model.POSTUserModel) (*model.POSTUserModel, error)
	FindAll(ctx context.Context) ([]*model.GETUserModel, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.GETUserModel, error)
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

	return postUserModelFromEntity(entity)
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.GETUserModel, error) {
	entity, err := r.client.User.Query().Where(user.ID(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return getUserModelFromEntity(entity)
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.GETUserModel, error) {
	entities, err := r.client.User.Query().Order(ent.Desc("created_at")).All(ctx)

	if err != nil {
		return nil, err
	}

	return userModelsFromEntities(entities)
}

func postUserModelFromEntity(entity *ent.User) (*model.POSTUserModel, error) {
	opts := []model.NewPOSTUserOption{
		model.NewPOSTUserID(entity.ID),
		model.NewPOSTUserFirstName(entity.FirstName),
		model.NewPOSTUserLastName(entity.LastName),
		model.NewPOSTUserEmail(entity.Email),
		model.NewPOSTUserPassword(entity.Password),
	}

	return model.NewPOSTUser(opts...)
}

func getUserModelFromEntity(entity *ent.User) (*model.GETUserModel, error) {
	opts := []model.NewGETUserOption{
		model.NewGETUserID(entity.ID),
		model.NewGETUserFirstName(entity.FirstName),
		model.NewGETUserLastName(entity.LastName),
		model.NewGETUserEmail(entity.Email),
	}

	return model.NewGETUser(opts...)
}

func userModelsFromEntities(entities []*ent.User) ([]*model.GETUserModel, error) {
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

	if errs[0] != nil {
		return nil, errs[0]
	}

	return results, nil
}
