package model

import "github.com/google/uuid"

type User struct {
	id         uuid.UUID
	first_name string
	last_name  string
	email      string
	password   string
}

type NewUserOption func(u *User)

func NewUser(opts ...NewUserOption) (*User, error) {
	user := &User{}

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

func NewUserID(id uuid.UUID) NewUserOption {
	return func(u *User) {
		u.id = id
	}
}

func NewUserFirstName(first_name string) NewUserOption {
	return func(u *User) {
		u.first_name = first_name
	}
}

func NewUserLastName(last_name string) NewUserOption {
	return func(u *User) {
		u.last_name = last_name
	}
}

func NewUserEmail(email string) NewUserOption {
	return func(u *User) {
		u.email = email
	}
}

func NewUserPassword(password string) NewUserOption {
	return func(u *User) {
		u.password = password
	}
}

func (user *User) GetID() uuid.UUID {
	return user.id
}
func (user *User) GetFirstName() string {
	return user.first_name
}
func (user *User) GetLastName() string {
	return user.last_name
}
func (user *User) GetEmail() string {
	return user.email
}
func (user *User) GetPassword() string {
	return user.password
}

func (u *User) validate() *ValidationErrors {
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

func (u *User) isIDValid() *ValidationError {
	if u.id == uuid.Nil {
		return NewValidationError("empty UUID in user ID is not allowed")
	}
	return nil
}
func (u *User) isUserFirstNameValid() *ValidationError {
	if u.first_name == "" {
		return NewValidationError("empty string in first name is not allowed")
	}
	return nil
}
func (u *User) isUserLastNameValid() *ValidationError {
	if u.last_name == "" {
		return NewValidationError("empty string in last name is not allowed")
	}
	return nil
}
func (u *User) isUserEmailValid() *ValidationError {
	if u.email == "" {
		return NewValidationError("empty string in email is not allowed")
	}
	return nil
}
func (u *User) isUserPasswordValid() *ValidationError {
	if u.password == "" {
		return NewValidationError("empty string in password is not allowed")
	}
	return nil
}
