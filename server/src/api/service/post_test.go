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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PostService_CreatePost(t *testing.T) {
	type args struct {
		ctx context.Context
		rb  request.POSTPostRequestBody
	}

	id := uuid.New()
	categoryId := rand.Int63n(1000) + 1
	userId := uuid.New()

	m1, err1 := model.NewPOSTPost(
		model.NewPOSTPostID(id),
		model.NewPOSTPostTitle("テストタイトル"),
		model.NewPOSTPostBody("テスト本文"),
		model.NewPOSTPostCategoryID(categoryId),
		model.NewPOSTPostUserID(userId),
	)
	require.NoError(t, err1)

	user1, ue1 := model.NewGETUser(
		model.NewGETUserID(userId),
		model.NewGETUserFirstName("悟"),
		model.NewGETUserLastName("五条"),
		model.NewGETUserEmail("s.gojo@kosen.com"),
	)
	require.NoError(t, ue1)

	category1, ce1 := model.NewPOSTCategoryAfterCreated(
		model.NewPOSTCategoryID(categoryId),
		model.NewPOSTCategoryName("テストカテゴリ"),
	)
	require.NoError(t, ce1)

	tests := []struct {
		name          string
		args          args
		want          *model.POSTPostModel
		prepareMockFn func(*testing.T, *postServiceMocks, *model.POSTPostModel, args)
		matcher       func(*testing.T, *model.POSTPostModel, *model.POSTPostModel, error)
	}{
		{
			name: "正常系：テスト記事を登録する",
			args: args{
				ctx: context.Background(),
				rb: request.POSTPostRequestBody{
					Title:      "テストタイトル",
					Body:       "テスト本文",
					CategoryID: categoryId,
					UserID:     userId,
				},
			},
			want: m1,
			prepareMockFn: func(t *testing.T, mocks *postServiceMocks, post *model.POSTPostModel, args args) {
				mocks.userRepo.EXPECT().FindByID(args.ctx, args.rb.UserID).Return(user1, nil)
				mocks.categoryRepo.EXPECT().FindByID(args.ctx, args.rb.CategoryID).Return(category1, nil)
				mocks.postRepo.EXPECT().Create(args.ctx,
					NewCmpMatcher(
						post,
						cmp.AllowUnexported(model.POSTPostModel{}),
						cmpopts.IgnoreFields(model.POSTPostModel{}, "id"),
					)).Return(post, nil)
			},
			matcher: func(t *testing.T, expected *model.POSTPostModel, got *model.POSTPostModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.POSTPostModel{}),
					cmpopts.IgnoreFields(model.POSTPostModel{}, "id"),
				}
				diff := cmp.Diff(expected, got, opts...)

				assert.NoError(t, err)
				assert.NotEmpty(t, got)
				assert.Empty(t, diff)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			service, mocks := getPostServiceMocks(ctrl)
			tt.prepareMockFn(t, mocks, tt.want, tt.args)

			// Act
			got, err := service.CreatePost(tt.args.ctx, tt.args.rb)

			// Assert
			tt.matcher(t, tt.want, got, err)
		})
	}
}

type postServiceMocks struct {
	postRepo     *mock_repository.MockPostRepository
	userRepo     *mock_repository.MockUserRepository
	categoryRepo *mock_repository.MockCategoryRepository
}

func getPostServiceMocks(ctrl *gomock.Controller) (*postService, *postServiceMocks) {
	mocks := &postServiceMocks{
		postRepo:     mock_repository.NewMockPostRepository(ctrl),
		userRepo:     mock_repository.NewMockUserRepository(ctrl),
		categoryRepo: mock_repository.NewMockCategoryRepository(ctrl),
	}

	service := &postService{
		postRepo:     mocks.postRepo,
		userRepo:     mocks.userRepo,
		categoryRepo: mocks.categoryRepo,
	}

	return service, mocks
}
