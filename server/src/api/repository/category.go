package repository

import (
	"blog/api/model"
	"blog/ent"
	"blog/ent/category"
	"context"
)

type CategoryRepository interface {
	Create(ctx context.Context, m *model.POSTCategoryModel) (*model.POSTCategoryModel, error)
	FindByID(ctx context.Context, id int64) (*model.POSTCategoryModel, error)
}

type categoryRepository struct {
	client *ent.Client
}

func NewCategoryRepository(client *ent.Client) CategoryRepository {
	return &categoryRepository{client}
}

func (r *categoryRepository) Create(ctx context.Context, m *model.POSTCategoryModel) (*model.POSTCategoryModel, error) {
	entity, err := r.client.Category.Create().
		SetName(m.GetName()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return categoryModelFromEntity(entity)
}

func (r *categoryRepository) FindByID(ctx context.Context, id int64) (*model.POSTCategoryModel, error) {
	entity, err := r.client.Category.Query().Where(category.ID(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return categoryModelFromEntity(entity)
}

func categoryModelFromEntity(entity *ent.Category) (*model.POSTCategoryModel, error) {
	opts := []model.NewPOSTCategoryOption{
		model.NewPOSTCategoryID(entity.ID),
		model.NewPOSTCategoryName(entity.Name),
	}

	return model.NewPOSTCategoryAfterCreated(opts...)
}
