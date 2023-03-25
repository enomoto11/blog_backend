package repository

import (
	"blog/api/model"
	"blog/ent"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UserRepository_Create(t *testing.T) {
	type args struct {
		m   *model.POSTUserModel
		ctx context.Context
	}

	id := uuid.New()
	ctx := context.Background()

	m1, err1 := model.NewPOSTUser(
		model.NewPOSTUserID(id),
		model.NewPOSTUserFirstName("悟"),
		model.NewPOSTUserLastName("五条"),
		model.NewPOSTUserEmail("s.gojo@gmail.com"),
		model.NewPOSTUserPassword("jujutsukaisenn"),
	)
	require.NoError(t, err1)

	tests := []struct {
		name    string
		args    args
		matcher func(t *testing.T, expect *model.POSTUserModel, got *model.POSTUserModel, err error)
	}{
		{
			name: "正常系：五条悟をユーザー登録する",
			args: args{
				m:   m1,
				ctx: ctx,
			},
			matcher: func(t *testing.T, expected *model.POSTUserModel, got *model.POSTUserModel, err error) {
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
			repo := &userRepository{client}

			// Action
			got, err := repo.Create(tt.args.ctx, tt.args.m)

			// Assert
			tt.matcher(t, tt.args.m, got, err)
		})
	}
}

func Test_UserRepository_FindAll(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		prepareFn func(t *testing.T, client *ent.Client) []*model.GETUserModel
		matcher   func(t *testing.T, expected []*model.GETUserModel, got []*model.GETUserModel, err error)
	}{
		{
			name: "正常系：ユーザーを全件取得する",
			prepareFn: func(t *testing.T, client *ent.Client) []*model.GETUserModel {
				var entities []*ent.User

				entity1, err := client.User.Create().
					SetID(uuid.New()).
					SetFirstName("悟").
					SetLastName("五条").
					SetEmail("satoru@kousen.com").
					SetPassword("jujutsukaisenn").
					Save(ctx)
				require.NoError(t, err)
				entities = append(entities, entity1)

				entity2, err := client.User.Create().
					SetID(uuid.New()).
					SetFirstName("空").
					SetLastName("五条").
					SetEmail("sora@daredayo.com").
					SetPassword("jujutsukaisenn").
					Save(ctx)
				require.NoError(t, err)
				entities = append(entities, entity2)

				models, _ := userModelsFromEntities(entities)

				return models
			},
			matcher: func(t *testing.T, expected []*model.GETUserModel, got []*model.GETUserModel, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
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
			repo := &userRepository{client}
			expected := tt.prepareFn(t, client)

			// Action
			got, err := repo.FindAll(ctx)

			// Assert
			tt.matcher(t, expected, got, err)
		})
	}
}
