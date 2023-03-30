package repository

import (
	"blog/api/model"
	"blog/ent"
	"blog/ent/post"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, m *model.PostModel) (*model.PostModel, error)
	FindAll(ctx context.Context) ([]*model.PostModel, error)
	FindByCategoryID(ctx context.Context, categoryID int64) ([]*model.PostModel, error)
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

func (r *postRepository) FindAll(ctx context.Context) ([]*model.PostModel, error) {
	entities, err := r.client.Post.Query().Order(ent.Desc("created_at")).All(ctx)
	if err != nil {
		return nil, err
	}

	var posts []*model.PostModel
	for _, entity := range entities {
		post, err := postModelFromEntity(entity)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *postRepository) FindByCategoryID(ctx context.Context, categoryID int64) ([]*model.PostModel, error) {
	entities, err := r.client.Post.Query().Where(post.CategoryID(categoryID)).All(ctx)
	if err != nil {
		return nil, err
	}

	var posts []*model.PostModel
	for _, entity := range entities {
		post, err := postModelFromEntity(entity)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
