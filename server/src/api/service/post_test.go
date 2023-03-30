package service

import (
	"blog/api/controller/request"
	"blog/api/model"
	mock_repository "blog/api/repository/mock"
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
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

	m1, err1 := model.NewPost(
		model.NewPostID(id),
		model.NewPostTitle("テストタイトル"),
		model.NewPostBody("テスト本文"),
		model.NewPostCategoryID(categoryId),
		model.NewPostUserID(userId),
	)
	require.NoError(t, err1)

	user1, ue1 := model.NewGETUser(
		model.NewGETUserID(userId),
		model.NewGETUserFirstName("悟"),
		model.NewGETUserLastName("五条"),
		model.NewGETUserEmail("s.gojo@kosen.com"),
	)
	require.NoError(t, ue1)

	category1, ce1 := model.NewCategoryAfterCreated(
		model.NewCategoryID(categoryId),
		model.NewCategoryName("テストカテゴリ"),
	)
	require.NoError(t, ce1)

	tests := []struct {
		name          string
		args          args
		want          *model.PostModel
		prepareMockFn func(*testing.T, *postServiceMocks, *model.PostModel, args)
		matcher       func(*testing.T, *model.PostModel, *model.PostModel, error)
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
			prepareMockFn: func(t *testing.T, mocks *postServiceMocks, post *model.PostModel, args args) {
				mocks.userRepo.EXPECT().FindByID(args.ctx, args.rb.UserID).Return(user1, nil)
				mocks.categoryRepo.EXPECT().FindByID(args.ctx, args.rb.CategoryID).Return(category1, nil)
				mocks.postRepo.EXPECT().Create(args.ctx,
					NewCmpMatcher(
						post,
						cmp.AllowUnexported(model.PostModel{}),
						cmpopts.IgnoreFields(model.PostModel{}, "id"),
					)).Return(post, nil)
			},
			matcher: func(t *testing.T, expected *model.PostModel, got *model.PostModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.PostModel{}),
					cmpopts.IgnoreFields(model.PostModel{}, "id"),
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

func Test_PostService_FindAllPosts(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	ctx := context.Background()

	id1 := uuid.New()
	id2 := uuid.New()
	categoryId := rand.Int63n(1000) + 1
	userId := uuid.New()

	m1, err1 := model.NewPost(
		model.NewPostID(id1),
		model.NewPostTitle("テストタイトル"),
		model.NewPostBody("テスト本文"),
		model.NewPostCategoryID(categoryId),
		model.NewPostUserID(userId),
	)
	require.NoError(t, err1)

	m2, err2 := model.NewPost(
		model.NewPostID(id2),
		model.NewPostTitle("テストタイトル2"),
		model.NewPostBody("テスト本文2"),
		model.NewPostCategoryID(categoryId),
		model.NewPostUserID(userId),
	)
	require.NoError(t, err2)

	tests := []struct {
		name          string
		args          args
		want          []*model.PostModel
		prepareMockFn func(*testing.T, *postServiceMocks, []*model.PostModel, args)
		matcher       func(*testing.T, []*model.PostModel, []*model.PostModel, error)
	}{
		{
			name: "正常系：テスト記事を全件取得する",
			args: args{
				ctx: ctx,
			},
			want: []*model.PostModel{m1, m2},
			prepareMockFn: func(t *testing.T, mocks *postServiceMocks, posts []*model.PostModel, args args) {
				mocks.postRepo.EXPECT().FindAll(args.ctx).Return(posts, nil)
			},
			matcher: func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.PostModel{}),
					cmpopts.IgnoreFields(model.PostModel{}, "id"),
				}
				diff := cmp.Diff(expected, got, opts...)

				assert.NoError(t, err)
				assert.NotEmpty(t, got)
				assert.Empty(t, diff)
			},
		},

		{
			name: "正常系：テスト記事を全件取得する（0件）",
			args: args{
				ctx: ctx,
			},
			want: nil,
			prepareMockFn: func(t *testing.T, mocks *postServiceMocks, posts []*model.PostModel, args args) {
				mocks.postRepo.EXPECT().FindAll(args.ctx).Return(posts, nil)
			},
			matcher: func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.PostModel{}),
					cmpopts.IgnoreFields(model.PostModel{}, "id"),
				}
				diff := cmp.Diff(expected, got, opts...)

				assert.NoError(t, err)
				assert.Empty(t, got)
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
			got, err := service.FindAllPosts(tt.args.ctx)

			// Assert
			tt.matcher(t, tt.want, got, err)
		})
	}
}

func Test_PostService_FindByCategoryID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	id1 := uuid.New()
	id2 := uuid.New()
	categoryId := rand.Int63n(1000) + 1
	userId := uuid.New()

	ctx := &gin.Context{}
	ctx.Params = []gin.Param{{Key: "id", Value: strconv.FormatInt(categoryId, 10)}}

	m1, err1 := model.NewPost(
		model.NewPostID(id1),
		model.NewPostTitle("テストタイトル"),
		model.NewPostBody("テスト本文"),
		model.NewPostCategoryID(categoryId),
		model.NewPostUserID(userId),
	)
	require.NoError(t, err1)

	m2, err2 := model.NewPost(
		model.NewPostID(id2),
		model.NewPostTitle("テストタイトル2"),
		model.NewPostBody("テスト本文2"),
		model.NewPostCategoryID(categoryId),
		model.NewPostUserID(userId),
	)
	require.NoError(t, err2)

	tests := []struct {
		name          string
		args          args
		want          []*model.PostModel
		prepareMockFn func(*testing.T, *postServiceMocks, []*model.PostModel, args)
		matcher       func(*testing.T, []*model.PostModel, []*model.PostModel, error)
	}{
		{
			name: "正常系：テスト記事をカテゴリIDで取得する",
			args: args{
				ctx: ctx,
			},
			want: []*model.PostModel{m1, m2},
			prepareMockFn: func(t *testing.T, mocks *postServiceMocks, posts []*model.PostModel, args args) {
				mocks.postRepo.EXPECT().FindByCategoryID(args.ctx, categoryId).Return(posts, nil)
			},
			matcher: func(t *testing.T, expected []*model.PostModel, got []*model.PostModel, err error) {
				// idは自動採番なので、比較対象から除外する
				opts := []cmp.Option{
					cmp.AllowUnexported(model.PostModel{}),
					cmpopts.IgnoreFields(model.PostModel{}, "id"),
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
			got, err := service.FindByCategoryID(tt.args.ctx)

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
