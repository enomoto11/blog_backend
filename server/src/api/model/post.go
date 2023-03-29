package model

import "github.com/google/uuid"

type PostModel struct {
	id          uuid.UUID
	title       string
	body        string
	user_id     uuid.UUID
	category_id int64
}

type NewPostOption func(p *PostModel)

func NewPost(opts ...NewPostOption) (*PostModel, error) {
	post := &PostModel{}

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

func NewPostID(id uuid.UUID) NewPostOption {
	return func(p *PostModel) {
		p.id = id
	}
}

func NewPostTitle(title string) NewPostOption {
	return func(p *PostModel) {
		p.title = title
	}
}

func NewPostBody(body string) NewPostOption {
	return func(p *PostModel) {
		p.body = body
	}
}

func NewPostUserID(user_id uuid.UUID) NewPostOption {
	return func(p *PostModel) {
		p.user_id = user_id
	}
}

func NewPostCategoryID(category_id int64) NewPostOption {
	return func(p *PostModel) {
		p.category_id = category_id
	}
}

func (post *PostModel) GetID() uuid.UUID {
	return post.id
}

func (post *PostModel) GetTitle() string {
	return post.title
}

func (post *PostModel) GetBody() string {
	return post.body
}

func (post *PostModel) GetUserID() uuid.UUID {
	return post.user_id
}

func (post *PostModel) GetCategoryID() int64 {
	return post.category_id
}

func (post *PostModel) validate() *ValidationErrors {
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

func (post *PostModel) isIDValid() *ValidationError {
	if post.id == uuid.Nil {
		return NewValidationError("empty UUID in post ID is not allowed")
	}
	return nil
}

func (post *PostModel) isTitleValid() *ValidationError {
	if post.title == "" {
		return NewValidationError("empty string in post title is not allowed")
	}
	return nil
}

func (post *PostModel) isBodyValid() *ValidationError {
	if post.body == "" {
		return NewValidationError("empty string in post body is not allowed")
	}
	return nil
}

func (post *PostModel) isUserIDValid() *ValidationError {
	if post.user_id == uuid.Nil {
		return NewValidationError("empty UUID in post user ID is not allowed")
	}
	return nil
}

func (post *PostModel) isCategoryIDValid() *ValidationError {
	if post.category_id == 0 {
		return NewValidationError("empty number in post category ID is not allowed")
	}
	return nil
}
