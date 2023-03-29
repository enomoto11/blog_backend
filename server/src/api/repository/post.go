package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, m *model.PostModel) (*model.PostModel, error)
}

type postRepository struct {
	client *ent.Client
}

func NewPostRepository(client *ent.Client) PostRepository {
	return &postRepository{client}
}

func (r *postRepository) Create(ctx context.Context, m *model.PostModel) (*model.PostModel, error) {
	entity, err := r.client.Post.Create().
		SetID(m.GetID()).
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

func postModelFromEntity(entity *ent.Post) (*model.PostModel, error) {
	opts := []model.NewPostOption{
		model.NewPostID(entity.ID),
		model.NewPostTitle(entity.Title),
		model.NewPostBody(entity.Body),
		model.NewPostCategoryID(entity.CategoryID),
		model.NewPostUserID(entity.UserID),
	}

	return model.NewPost(opts...)
}
