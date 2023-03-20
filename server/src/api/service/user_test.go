package service

import (
	"blog/api/controller/request"
	"blog/api/model"
	mock_repository "blog/api/repository/mock"
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
		requestBody   request.CreateUserRequestBody
		InitModel     func(*testing.T, request.CreateUserRequestBody) *model.User
		prepareMockFn func(*testing.T, *userServiceMocks, *model.User)
		matcher       func(*testing.T, *model.User, *model.User, error)
	}{
		{
			name: "正常系",
			requestBody: request.CreateUserRequestBody{
				First_name: "アキ",
				Last_name:  "早川",
				Email:      "a.hayakawa@koan.me",
				Password:   "kon_mirai4",
			},
			InitModel: func(t *testing.T, rb request.CreateUserRequestBody) *model.User {
				m, err := model.NewUser(
					model.NewUserID(id),
					model.NewUserFirstName(rb.First_name),
					model.NewUserLastName(rb.Last_name),
					model.NewUserEmail(rb.Email),
					model.NewUserPassword(rb.Password),
				)
				require.NoError(t, err)

				return m
			},
			prepareMockFn: func(t *testing.T, mocks *userServiceMocks, user *model.User) {
				mocks.createUserRepo.EXPECT().Create(ctx,
					NewCmpMatcher(
						user,
						cmp.AllowUnexported(model.User{}),
						cmpopts.IgnoreFields(model.User{}, "id"),
					),
				).Return(user, nil)
			},
			matcher: func(t *testing.T, expected *model.User, got *model.User, err error) {
				opts := []cmp.Option{
					cmp.AllowUnexported(model.User{}),
					cmpopts.IgnoreFields(model.User{}, "id"),
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
