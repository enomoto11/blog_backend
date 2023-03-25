package model

import "fmt"

const (
	postCategoryNameMaxLength = 20
)

type POSTCategoryModel struct {
	id   int
	name string
}

type NewPOSTCategoryOption func(c *POSTCategoryModel)

func NewPOSTCategory(opts ...NewPOSTCategoryOption) (*POSTCategoryModel, error) {
	category := &POSTCategoryModel{}

	for _, opt := range opts {
		opt(category)
	}
	if err := category.validate(); err != nil {
		return nil, NewValidationError(err.Error())
	}

	return category, nil
}

func NewPOSTCategoryID(id int) NewPOSTCategoryOption {
	return func(c *POSTCategoryModel) {
		c.id = id
	}
}

func NewPOSTCategoryName(name string) NewPOSTCategoryOption {
	return func(c *POSTCategoryModel) {
		c.name = name
	}
}

func (category *POSTCategoryModel) GetID() int {
	return category.id
}
func (category *POSTCategoryModel) GetName() string {
	return category.name
}

func (category *POSTCategoryModel) validate() *ValidationErrors {
	var errors []*ValidationError

	if ve := category.isIDValid(); ve != nil {
		errors = append(errors, ve)
	}
	if ve := category.isNameValid(); ve != nil {
		errors = append(errors, ve)
	}

	return validationErrorSliceToValidationErrors(errors)
}

func (category *POSTCategoryModel) isIDValid() *ValidationError {
	if category.id <= 0 {
		return NewValidationError("category id must be greater than 0")
	}
	return nil
}
func (category *POSTCategoryModel) isNameValid() *ValidationError {
	if category.name == "" {
		return NewValidationError("empty string in category name is not allowed")
	}
	if len(category.name) > postCategoryNameMaxLength {
		return NewValidationError(fmt.Sprintf("category name must be less than %d characters", postCategoryNameMaxLength))
	}
	return nil
}
