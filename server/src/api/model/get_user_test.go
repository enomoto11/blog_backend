package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_UserModel_NewUser(t *testing.T) {
	type args struct {
		options []NewGETUserOption
	}
	id := uuid.New()

	tests := []struct {
		name    string
		args    args
		want    *GETUserModel
		wantErr string
	}{
		{
			name: "正常系",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(id),
					NewGETUserFirstName("桜木"),
					NewGETUserLastName("花道"),
					NewGETUserEmail("h.sakuragi@shohoku.high.com"),
				},
			},
			want: &GETUserModel{
				id:         id,
				first_name: "桜木",
				last_name:  "花道",
				email:      "h.sakuragi@shohoku.high.com",
			},
		},
		{
			name: "異常系: idがnilの時エラーが返る",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(uuid.Nil),
					NewGETUserFirstName("桜木"),
					NewGETUserLastName("花道"),
					NewGETUserEmail("h.sakuragi@shohoku.high.com"),
				},
			},
			wantErr: "empty UUID in user ID is not allowed",
		},
		{
			name: "異常系: first_name が空文字の時エラーが返る",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(id),
					NewGETUserFirstName(""),
					NewGETUserLastName("花道"),
					NewGETUserEmail("h.sakuragi@shohoku.high.com"),
				},
			},
			wantErr: "empty string in first name is not allowed",
		},
		{
			name: "異常系: last_name が空文字の時エラーが返る",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(id),
					NewGETUserFirstName("桜木"),
					NewGETUserLastName(""),
					NewGETUserEmail("h.sakuragi@shohoku.high.com"),
				},
			},
			wantErr: "empty string in last name is not allowed",
		},
		{
			name: "異常系: email が空文字の時エラーが返る",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(id),
					NewGETUserFirstName("桜木"),
					NewGETUserLastName("花道"),
					NewGETUserEmail(""),
				},
			},
			wantErr: "empty string in email is not allowed",
		},
		{
			name: "異常系: email がformatに則っていない時エラーが返る",
			args: args{
				[]NewGETUserOption{
					NewGETUserID(id),
					NewGETUserFirstName("桜木"),
					NewGETUserLastName("花道"),
					NewGETUserEmail("mail-address"),
				},
			},
			wantErr: "invalid email format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGETUser(tt.args.options...)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
