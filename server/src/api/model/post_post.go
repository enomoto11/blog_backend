package model

import "github.com/google/uuid"

type POSTPostModel struct {
	id          uuid.UUID
	title       string
	body        string
	user_id     uuid.UUID
	category_id int64
}

type NewPOSTPostOption func(p *POSTPostModel)

func NewPOSTPost(opts ...NewPOSTPostOption) (*POSTPostModel, error) {
	post := &POSTPostModel{}

	// post作成時にデフォルトでuuidを設定。上書き可能。
	post.id = uuid.Must(uuid.NewRandom())

	for _, opt := range opts {
		opt(post)
	}
	if err := post.validate(); err != nil {
		return nil, NewValidationError(err.Error())
	}

	return post, nil
}

func NewPOSTPostID(id uuid.UUID) NewPOSTPostOption {
	return func(p *POSTPostModel) {
		p.id = id
	}
}

func NewPOSTPostTitle(title string) NewPOSTPostOption {
	return func(p *POSTPostModel) {
		p.title = title
	}
}

func NewPOSTPostBody(body string) NewPOSTPostOption {
	return func(p *POSTPostModel) {
		p.body = body
	}
}

func NewPOSTPostUserID(user_id uuid.UUID) NewPOSTPostOption {
	return func(p *POSTPostModel) {
		p.user_id = user_id
	}
}

func NewPOSTPostCategoryID(category_id int64) NewPOSTPostOption {
	return func(p *POSTPostModel) {
		p.category_id = category_id
	}
}

func (post *POSTPostModel) GetID() uuid.UUID {
	return post.id
}

func (post *POSTPostModel) GetTitle() string {
	return post.title
}

func (post *POSTPostModel) GetBody() string {
	return post.body
}

func (post *POSTPostModel) GetUserID() uuid.UUID {
	return post.user_id
}

func (post *POSTPostModel) GetCategoryID() int64 {
	return post.category_id
}

func (post *POSTPostModel) validate() *ValidationErrors {
	var errors []*ValidationError
	if err := post.isIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := post.isTitleValid(); err != nil {
		errors = append(errors, err)
	}
	if err := post.isBodyValid(); err != nil {
		errors = append(errors, err)
	}
	if err := post.isUserIDValid(); err != nil {
		errors = append(errors, err)
	}
	if err := post.isCategoryIDValid(); err != nil {
		errors = append(errors, err)
	}

	return validationErrorSliceToValidationErrors(errors)
}

func (post *POSTPostModel) isIDValid() *ValidationError {
	if post.id == uuid.Nil {
		return NewValidationError("empty UUID in post ID is not allowed")
	}
	return nil
}

func (post *POSTPostModel) isTitleValid() *ValidationError {
	if post.title == "" {
		return NewValidationError("empty string in post title is not allowed")
	}
	return nil
}

func (post *POSTPostModel) isBodyValid() *ValidationError {
	if post.body == "" {
		return NewValidationError("empty string in post body is not allowed")
	}
	return nil
}

func (post *POSTPostModel) isUserIDValid() *ValidationError {
	if post.user_id == uuid.Nil {
		return NewValidationError("empty UUID in post user ID is not allowed")
	}
	return nil
}

func (post *POSTPostModel) isCategoryIDValid() *ValidationError {
	if post.category_id == 0 {
		return NewValidationError("empty number in post category ID is not allowed")
	}
	return nil
}
