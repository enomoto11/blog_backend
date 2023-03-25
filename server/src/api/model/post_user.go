package model

import (
	"regexp"

	"github.com/google/uuid"
)

type POSTUserModel struct {
	id         uuid.UUID
	first_name string
	last_name  string
	email      string
	password   string
}

type NewPOSTUserOption func(u *POSTUserModel)

func NewPOSTUser(opts ...NewPOSTUserOption) (*POSTUserModel, error) {
	user := &POSTUserModel{}

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

func NewPOSTUserID(id uuid.UUID) NewPOSTUserOption {
	return func(u *POSTUserModel) {
		u.id = id
	}
}

func NewPOSTUserFirstName(first_name string) NewPOSTUserOption {
	return func(u *POSTUserModel) {
		u.first_name = first_name
	}
}

func NewPOSTUserLastName(last_name string) NewPOSTUserOption {
	return func(u *POSTUserModel) {
		u.last_name = last_name
	}
}

func NewPOSTUserEmail(email string) NewPOSTUserOption {
	return func(u *POSTUserModel) {
		u.email = email
	}
}

func NewPOSTUserPassword(password string) NewPOSTUserOption {
	return func(u *POSTUserModel) {
		u.password = password
	}
}

func (user *POSTUserModel) GetID() uuid.UUID {
	return user.id
}
func (user *POSTUserModel) GetFirstName() string {
	return user.first_name
}
func (user *POSTUserModel) GetLastName() string {
	return user.last_name
}
func (user *POSTUserModel) GetEmail() string {
	return user.email
}
func (user *POSTUserModel) GetPassword() string {
	return user.password
}

func (u *POSTUserModel) validate() *ValidationErrors {
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
	if ve := u.isUserPasswordValid(); ve != nil {
		errors = append(errors, ve)
	}

	return validationErrorSliceToValidationErrors(errors)

}

func (u *POSTUserModel) isIDValid() *ValidationError {
	if u.id == uuid.Nil {
		return NewValidationError("empty UUID in user ID is not allowed")
	}
	return nil
}
func (u *POSTUserModel) isUserFirstNameValid() *ValidationError {
	if u.first_name == "" {
		return NewValidationError("empty string in first name is not allowed")
	}
	return nil
}
func (u *POSTUserModel) isUserLastNameValid() *ValidationError {
	if u.last_name == "" {
		return NewValidationError("empty string in last name is not allowed")
	}
	return nil
}
func (u *POSTUserModel) isUserEmailValid() *ValidationError {
	if u.email == "" {
		return NewValidationError("empty string in email is not allowed")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.email) {
		return NewValidationError("invalid email format")
	}
	return nil
}
func (u *POSTUserModel) isUserPasswordValid() *ValidationError {
	if u.password == "" {
		return NewValidationError("empty string in password is not allowed")
	}
	return nil
}
