package model

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CategoryModel_NewCategory(t *testing.T) {
	type args struct {
		options []NewPOSTCategoryOption
	}

	id := rand.Int63n(1000) * 1

	tests := []struct {
		name    string
		args    args
		want    *POSTCategoryModel
		wantErr string
	}{
		{
			name: "正常系",
			args: args{
				[]NewPOSTCategoryOption{
					NewPOSTCategoryID(id),
					NewPOSTCategoryName("テストカテゴリー"),
				},
			},
			want: &POSTCategoryModel{
				id:   id,
				name: "テストカテゴリー",
			},
		},
		{
			name: "異常系: nameが空文字の時エラーが返る",
			args: args{
				[]NewPOSTCategoryOption{
					NewPOSTCategoryID(id),
					NewPOSTCategoryName(""),
				},
			},
			wantErr: "empty string in category name is not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, err1 := NewPOSTCategoryBeforeCreated(tt.args.options...)
			got2, err2 := NewPOSTCategoryAfterCreated(tt.args.options...)
			if tt.wantErr != "" {
				assert.EqualError(t, err1, tt.wantErr)
			} else {
				assert.NoError(t, err1)
				assert.NoError(t, err2)
				assert.Equal(t, tt.want, got1)
				assert.Equal(t, tt.want, got2)
			}
		})
	}
}
