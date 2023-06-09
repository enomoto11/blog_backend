package service

import (
	"blog/api/controller/request"
	"blog/api/model"
	mock_repository "blog/api/repository/mock"
	"context"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CategoryService_CreateCategory(t *testing.T) {
	type args struct {
		ctx context.Context
		rb  request.POSTCategoryRequestBody
	}

	id := rand.Int63n(1000) + 1

	m1, err1 := model.NewCategoryAfterCreated(
		model.NewCategoryID(id),
		model.NewCategoryName("テストカテゴリー"),
	)
	require.NoError(t, err1)

	tests := []struct {
		name          string
		args          args
		want          *model.CategoryModel
		prepareMockFn func(*testing.T, *categoryServiceMocks, *model.CategoryModel, args)
		matcher       func(*testing.T, *model.CategoryModel, *model.CategoryModel, error)
	}{
		{
			name: "正常系：テストカテゴリーを登録する",
			args: args{
				ctx: context.Background(),
				rb: request.POSTCategoryRequestBody{
					Name: "テストカテゴリー",
				},
			},
			want: m1,
			prepareMockFn: func(t *testing.T, mocks *categoryServiceMocks, category *model.CategoryModel, args args) {
				mocks.categoryRepo.EXPECT().Create(args.ctx,
					NewCmpMatcher(
						category,
						cmp.AllowUnexported(model.CategoryModel{}),
						cmpopts.IgnoreFields(model.CategoryModel{}, "id"),
					),
				).Return(category, nil)
			},
			matcher: func(t *testing.T, expected *model.CategoryModel, got *model.CategoryModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.CategoryModel{}),
					cmpopts.IgnoreFields(model.CategoryModel{}, "id"),
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service, mocks := getCategoryServiceMocks(ctrl)
			tt.prepareMockFn(t, mocks, tt.want, tt.args)

			// Act
			got, err := service.CreateCategory(tt.args.ctx, tt.args.rb)

			// Assert
			tt.matcher(t, tt.want, got, err)
		})
	}

}

type categoryServiceMocks struct {
	categoryRepo *mock_repository.MockCategoryRepository
}

func getCategoryServiceMocks(ctrl *gomock.Controller) (*categoryService, *categoryServiceMocks) {
	mocks := &categoryServiceMocks{
		categoryRepo: mock_repository.NewMockCategoryRepository(ctrl),
	}

	service := &categoryService{
		categoryRepo: mocks.categoryRepo,
	}

	return service, mocks
}
