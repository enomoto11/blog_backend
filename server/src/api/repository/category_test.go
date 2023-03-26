package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CategoryRepository_Create(t *testing.T) {
	type args struct {
		m   *model.POSTCategoryModel
		ctx context.Context
	}

	ctx := context.Background()

	m1, err1 := model.NewPOSTCategoryBeforeCreated(
		model.NewPOSTCategoryName("テストカテゴリー"),
	)

	require.NoError(t, err1)

	tests := []struct {
		name    string
		args    args
		matcher func(t *testing.T, expected *model.POSTCategoryModel, got *model.POSTCategoryModel, err error)
	}{
		{
			name: "正常系：テストカテゴリーを登録する",
			args: args{
				m:   m1,
				ctx: ctx,
			},
			matcher: func(t *testing.T, expected *model.POSTCategoryModel, got *model.POSTCategoryModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.POSTCategoryModel{}),
					cmpopts.IgnoreFields(model.POSTCategoryModel{}, "id"),
				}
				diff := cmp.Diff(expected, got, opts...)

				assert.NoError(t, err)
				assert.NotEmpty(t, got)
				assert.Equal(t, diff, "")
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
			repo := &categoryRepository{client}

			// Action
			got, err := repo.Create(tt.args.ctx, tt.args.m)

			// Assert
			tt.matcher(t, tt.args.m, got, err)
		})
	}
}

func Test_CategoryRepository_FindByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	ctx := context.Background()
	id := rand.Int63n(1000) + 1

	m1, err1 := model.NewPOSTCategoryAfterCreated(
		model.NewPOSTCategoryID(id),
		model.NewPOSTCategoryName("テストカテゴリー"),
	)
	require.NoError(t, err1)

	tests := []struct {
		name    string
		args    args
		matcher func(t *testing.T, expected *model.POSTCategoryModel, got *model.POSTCategoryModel, err error)

		// 事前に登録するデータ
		beforeCreate func(t *testing.T, ctx context.Context, client *ent.Client)

		// 事後に削除するデータ
		afterDelete func(t *testing.T, ctx context.Context, client *ent.Client)
	}{
		{
			name: "正常系：テストカテゴリーを取得する",
			args: args{
				ctx: ctx,
				id:  id,
			},
			matcher: func(t *testing.T, expected *model.POSTCategoryModel, got *model.POSTCategoryModel, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
				assert.Equal(t, expected, got)
			},
			beforeCreate: func(t *testing.T, ctx context.Context, client *ent.Client) {
				_, err := client.Category.Create().SetID(id).SetName("テストカテゴリー").Save(ctx)
				require.NoError(t, err)
			},
			afterDelete: func(t *testing.T, ctx context.Context, client *ent.Client) {
				err := client.Category.DeleteOneID(id).Exec(ctx)
				require.NoError(t, err)
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
			repo := &categoryRepository{client}

			// 事前に登録するデータ
			if tt.beforeCreate != nil {
				tt.beforeCreate(t, tt.args.ctx, client)
			}

			// Action
			got, err := repo.FindByID(tt.args.ctx, tt.args.id)

			// Assert
			tt.matcher(t, m1, got, err)

			// 事後に削除するデータ
			if tt.afterDelete != nil {
				tt.afterDelete(t, tt.args.ctx, client)
			}
		})
	}
}
