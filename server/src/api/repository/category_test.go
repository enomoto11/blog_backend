package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
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
