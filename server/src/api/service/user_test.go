package service

import (
	"blog/api/controller/request"
	"blog/api/model"
	mock_repository "blog/api/repository/mock"
	"blog/ent"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UserService_CreateUser(t *testing.T) {
	id := uuid.New()
	ctx := context.Background()

	tests := []struct {
		name          string
		requestBody   request.POSTUserRequestBody
		InitModel     func(*testing.T, request.POSTUserRequestBody) *model.POSTUserModel
		prepareMockFn func(*testing.T, *userServiceMocks, *model.POSTUserModel)
		matcher       func(*testing.T, *model.POSTUserModel, *model.POSTUserModel, error)
	}{
		{
			name: "正常系",
			requestBody: request.POSTUserRequestBody{
				FirstName: "アキ",
				LastName:  "早川",
				Email:     "a.hayakawa@koan.me",
				Password:  "kon_mirai4",
			},
			InitModel: func(t *testing.T, rb request.POSTUserRequestBody) *model.POSTUserModel {
				m, err := model.NewPOSTUser(
					model.NewPOSTUserID(id),
					model.NewPOSTUserFirstName(rb.FirstName),
					model.NewPOSTUserLastName(rb.LastName),
					model.NewPOSTUserEmail(rb.Email),
					model.NewPOSTUserPassword(rb.Password),
				)
				require.NoError(t, err)

				return m
			},
			prepareMockFn: func(t *testing.T, mocks *userServiceMocks, user *model.POSTUserModel) {
				mocks.createUserRepo.EXPECT().Create(ctx,
					NewCmpMatcher(
						user,
						cmp.AllowUnexported(model.POSTUserModel{}),
						cmpopts.IgnoreFields(model.POSTUserModel{}, "id"),
					),
				).Return(user, nil)
			},
			matcher: func(t *testing.T, expected *model.POSTUserModel, got *model.POSTUserModel, err error) {
				opts := []cmp.Option{
					cmp.AllowUnexported(model.POSTUserModel{}),
					cmpopts.IgnoreFields(model.POSTUserModel{}, "id"),
				}
				diff := cmp.Diff(expected, got, opts...)

				assert.Nil(t, err)
				assert.Equal(t, diff, "")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc, mocks := getServiceMocks(ctrl)
			expected := tt.InitModel(t, tt.requestBody)
			tt.prepareMockFn(t, mocks, expected)

			// Action
			got, err := svc.CreateUser(ctx, tt.requestBody)

			// Assert
			tt.matcher(t, expected, got, err)
		})
	}
}

func Test_UserService_FindAllUsers(t *testing.T) {
	ctx := context.Background()

	id1 := uuid.New()
	id2 := uuid.New()

	m1, err1 := model.NewGETUser(
		model.NewGETUserID(id1),
		model.NewGETUserFirstName("アキ"),
		model.NewGETUserLastName("早川"),
		model.NewGETUserEmail("a.hayakawa@koan.com"),
	)
	require.NoError(t, err1)

	m2, err2 := model.NewGETUser(
		model.NewGETUserID(id2),
		model.NewGETUserFirstName("デンジ"),
		model.NewGETUserLastName("チェーンソーマン"),
		model.NewGETUserEmail("d.chainsaw@koan.com"),
	)
	require.NoError(t, err2)

	tests := []struct {
		name          string
		prepareFn     func(*testing.T, *ent.Client)
		prepareMockFn func(*testing.T, *userServiceMocks)
		matcher       func(*testing.T, []*model.GETUserModel, error)
	}{
		{
			name: "正常系",
			prepareFn: func(t *testing.T, client *ent.Client) {
				_, err1 := client.User.Create().
					SetID(id1).
					SetFirstName("アキ").
					SetLastName("早川").
					SetEmail("a.hayakawa@koan.com").
					SetPassword("kon_mirai4").
					Save(ctx)
				require.NoError(t, err1)

				_, err2 := client.User.Create().
					SetID(id2).
					SetFirstName("デンジ").
					SetLastName("チェーンソーマン").
					SetEmail("d.chainsaw@koan.com").
					SetPassword("makima_san").
					Save(ctx)
				require.NoError(t, err2)
			},
			prepareMockFn: func(t *testing.T, mocks *userServiceMocks) {

				mocks.createUserRepo.EXPECT().FindAll(ctx).Return([]*model.GETUserModel{
					m1,
					m2,
				}, nil)
			},
			matcher: func(t *testing.T, got []*model.GETUserModel, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, got)
				assert.Equal(t, []*model.GETUserModel{
					m1,
					m2,
				}, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc, mocks := getServiceMocks(ctrl)
			tt.prepareMockFn(t, mocks)

			// Action
			got, err := svc.FindAllUsers(ctx)

			// Assert
			tt.matcher(t, got, err)
		})
	}
}

type userServiceMocks struct {
	createUserRepo *mock_repository.MockUserRepository
}

func getServiceMocks(ctrl *gomock.Controller) (*userService, *userServiceMocks) {
	mocks := userServiceMocks{
		mock_repository.NewMockUserRepository(ctrl),
	}
	svc := userService{
		mocks.createUserRepo,
	}

	return &svc, &mocks
}
