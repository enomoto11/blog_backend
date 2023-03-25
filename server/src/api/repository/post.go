package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, m *model.POSTPostModel) (*model.POSTPostModel, error)
}

type postRepository struct {
	client *ent.Client
}

func NewPostRepository(client *ent.Client) PostRepository {
	return &postRepository{client}
}

func (r *postRepository) Create(ctx context.Context, m *model.POSTPostModel) (*model.POSTPostModel, error) {
	entity, err := r.client.Post.Create().
		SetTitle(m.GetTitle()).
		SetBody(m.GetBody()).
		SetCategoryID(m.GetCategoryID()).
		SetUserID(m.GetUserID()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return postModelFromEntity(entity)
}

func postModelFromEntity(entity *ent.Post) (*model.POSTPostModel, error) {
	opts := []model.NewPOSTPostOption{
		model.NewPOSTPostID(entity.ID),
		model.NewPOSTPostTitle(entity.Title),
		model.NewPOSTPostBody(entity.Body),
		model.NewPOSTPostCategoryID(entity.CategoryID),
		model.NewPOSTPostUserID(entity.UserID),
	}

	return model.NewPOSTPost(opts...)
}
