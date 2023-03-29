package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PosrRepository_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		m   *model.PostModel
	}

	id := uuid.New()
	userId := uuid.New()
	categoryId := rand.Int63n(1000) + 1

	m1, err1 := model.NewPost(
		model.NewPostID(id),
		model.NewPostTitle("テストタイトル"),
		model.NewPostBody("テスト本文"),
		model.NewPostUserID(userId),
		model.NewPostCategoryID(categoryId),
	)
	require.NoError(t, err1)

	tests := []struct {
		name      string
		args      args
		prepareFn func(t *testing.T, client *ent.Client, args args)
		matcher   func(t *testing.T, expected *model.PostModel, got *model.PostModel, err error)
	}{
		{
			name: "正常系：テスト記事を登録する",
			args: args{
				ctx: context.Background(),
				m:   m1,
			},
			prepareFn: func(t *testing.T, client *ent.Client, args args) {
				_, err := client.User.Create().
					SetID(args.m.GetUserID()).
					SetFirstName("悟").
					SetLastName("五条").
					SetEmail("s.gojo@kosen.com").
					SetPassword("jujutsukaisenn").
					Save(args.ctx)
				require.NoError(t, err)

				_, err = client.Category.Create().
					SetID(args.m.GetCategoryID()).
					SetName("テストカテゴリ").
					Save(args.ctx)
				require.NoError(t, err)
			},
			matcher: func(t *testing.T, expected *model.PostModel, got *model.PostModel, err error) {
				assert.NoError(t, err)
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arange
			client, err := InitializeEntClient(t)
			defer func(client *ent.Client) {
				_ = client.Close()
			}(client)
			require.NoError(t, err)
			repo := &postRepository{client}
			tt.prepareFn(t, client, tt.args)

			// Action
			got, err := repo.Create(tt.args.ctx, tt.args.m)

			// Assert
			tt.matcher(t, tt.args.m, got, err)
		})
	}
}

func Test_PosrRepository_FindAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	id1 := uuid.New()
	id2 := uuid.New()
	categoryID := rand.Int63n(1000) + 1
	userID := uuid.New()

	p1, err1 := model.NewPost(
		model.NewPostID(id1),
		model.NewPostTitle("テストタイトル"),
		model.NewPostBody("テスト本文"),
		model.NewPostUserID(userID),
		model.NewPostCategoryID(categoryID),
	)
	require.NoError(t, err1)

	p2, err2 := model.NewPost(
		model.NewPostID(id2),
		model.NewPostTitle("テストタイトル2"),
		model.NewPostBody("テスト本文2"),
		model.NewPostUserID(userID),
		model.NewPostCategoryID(categoryID),
	)
	require.NoError(t, err2)

	tests := []struct {
		name      string
		args      args
		expected  []*model.PostModel
		prepareFn func(t *testing.T, client *ent.Client, args args)
		matcher   func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error)
	}{
		{
			name: "正常系：テスト記事を全件取得する",
			args: args{
				ctx: context.Background(),
			},
			expected: []*model.PostModel{
				p1,
				p2,
			},
			prepareFn: func(t *testing.T, client *ent.Client, args args) {
				_, err := client.User.Create().
					SetID(userID).
					SetFirstName("ゴン").
					SetLastName("フリークス").
					SetEmail("f.gon@hunter.com").
					SetPassword("hunterhunter").
					Save(args.ctx)
				require.NoError(t, err)

				_, err = client.Category.Create().
					SetID(categoryID).
					SetName("テストカテゴリ").
					Save(args.ctx)
				require.NoError(t, err)

				_, err = client.Post.Create().
					SetID(id1).
					SetTitle("テストタイトル").
					SetBody("テスト本文").
					SetUserID(userID).
					SetCategoryID(categoryID).
					Save(args.ctx)
				require.NoError(t, err)

				_, err = client.Post.Create().
					SetID(id2).
					SetTitle("テストタイトル2").
					SetBody("テスト本文2").
					SetUserID(userID).
					SetCategoryID(categoryID).
					Save(args.ctx)
				require.NoError(t, err)
			},

			matcher: func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error) {
				assert.NoError(t, err)
				assert.Equal(t, expected, got)
			},
		},
		{
			name: "正常系：テスト記事を全件取得する（記事が存在しない）",
			args: args{
				ctx: context.Background(),
			},
			expected: nil,
			prepareFn: func(t *testing.T, client *ent.Client, args args) {
				_, err := client.User.Create().
					SetID(userID).
					SetFirstName("ゴン").
					SetLastName("フリークス").
					SetEmail("f.gon@hunter.com").
					SetPassword("hunterhunter").
					Save(args.ctx)
				require.NoError(t, err)

				_, err = client.Category.Create().
					SetID(categoryID).
					SetName("テストカテゴリ").
					Save(args.ctx)
				require.NoError(t, err)
			},
			matcher: func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error) {
				assert.NoError(t, err)
				assert.Equal(t, expected, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arange
			client, err := InitializeEntClient(t)
			defer func(client *ent.Client) {
				_ = client.Close()
			}(client)
			require.NoError(t, err)
			repo := &postRepository{client}
			tt.prepareFn(t, client, tt.args)

			// Action
			got, err := repo.FindAll(tt.args.ctx)

			// Assert
			tt.matcher(t, tt.expected, got, err)
		})
	}
}
