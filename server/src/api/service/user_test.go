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
		InitModel     func(*testing.T, request.CreateUserRequestBody) *model.POSTUserModel
		prepareMockFn func(*testing.T, *userServiceMocks, *model.POSTUserModel)
		matcher       func(*testing.T, *model.POSTUserModel, *model.POSTUserModel, error)
	}{
		{
			name: "正常系",
			requestBody: request.CreateUserRequestBody{
				First_name: "アキ",
				Last_name:  "早川",
				Email:      "a.hayakawa@koan.me",
				Password:   "kon_mirai4",
			},
			InitModel: func(t *testing.T, rb request.CreateUserRequestBody) *model.POSTUserModel {
				m, err := model.NewPOSTUser(
					model.NewPOSTUserID(id),
					model.NewPOSTUserFirstName(rb.First_name),
					model.NewPOSTUserLastName(rb.Last_name),
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
