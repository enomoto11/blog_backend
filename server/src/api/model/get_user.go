package model

import (
	"regexp"

	"github.com/google/uuid"
)

type GETUserModel struct {
	id         uuid.UUID
	first_name string
	last_name  string
	email      string
}

type NewGETUserOption func(u *GETUserModel)

func NewGETUser(opts ...NewGETUserOption) (*GETUserModel, error) {
	user := &GETUserModel{}

	// user作成時にデフォルトでuuidを設定。上書き可能。
	user.id = uuid.Must(uuid.NewRandom())

	for _, opt := range opts {
		opt(user)
	}
	if err := user.validate(); err != nil {
		return nil, NewValidationError(err.Error())
	}

	return user, nil
}

func NewGETUserID(id uuid.UUID) NewGETUserOption {
	return func(u *GETUserModel) {
		u.id = id
	}
}

func NewGETUserFirstName(first_name string) NewGETUserOption {
	return func(u *GETUserModel) {
		u.first_name = first_name
	}
}

func NewGETUserLastName(last_name string) NewGETUserOption {
	return func(u *GETUserModel) {
		u.last_name = last_name
	}
}

func NewGETUserEmail(email string) NewGETUserOption {
	return func(u *GETUserModel) {
		u.email = email
	}
}

func (user *GETUserModel) GetID() uuid.UUID {
	return user.id
}
func (user *GETUserModel) GetFirstName() string {
	return user.first_name
}
func (user *GETUserModel) GetLastName() string {
	return user.last_name
}
func (user *GETUserModel) GetEmail() string {
	return user.email
}

func (u *GETUserModel) validate() *ValidationErrors {
	var errors []*ValidationError

	if ve := u.isIDValid(); ve != nil {
		errors = append(errors, ve)
	}
	if ve := u.isUserFirstNameValid(); ve != nil {
		errors = append(errors, ve)
	}
	if ve := u.isUserLastNameValid(); ve != nil {
		errors = append(errors, ve)
	}
	if ve := u.isUserEmailValid(); ve != nil {
		errors = append(errors, ve)
	}

	return validationErrorSliceToValidationErrors(errors)

}

func (u *GETUserModel) isIDValid() *ValidationError {
	if u.id == uuid.Nil {
		return NewValidationError("empty UUID in user ID is not allowed")
	}
	return nil
}
func (u *GETUserModel) isUserFirstNameValid() *ValidationError {
	if u.first_name == "" {
		return NewValidationError("empty string in first name is not allowed")
	}
	return nil
}
func (u *GETUserModel) isUserLastNameValid() *ValidationError {
	if u.last_name == "" {
		return NewValidationError("empty string in last name is not allowed")
	}
	return nil
}
func (u *GETUserModel) isUserEmailValid() *ValidationError {
	if u.email == "" {
		return NewValidationError("empty string in email is not allowed")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.email) {
		return NewValidationError("invalid email format")
	}
	return nil
}
