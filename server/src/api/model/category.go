package model

import (
	"fmt"
)

const (
	categoryNameMaxLength = 50
)

type CategoryModel struct {
	id   int64
	name string
}

type NewCategoryOption func(c *CategoryModel)

func NewCategoryBeforeCreated(opts ...NewCategoryOption) (*CategoryModel, error) {
	category := &CategoryModel{}

	for _, opt := range opts {
		opt(category)
	}
	if err := category.validate(false); err != nil {
		return nil, NewValidationError(err.Error())
	}

	return category, nil
}

func NewCategoryAfterCreated(opts ...NewCategoryOption) (*CategoryModel, error) {
	category := &CategoryModel{}

	for _, opt := range opts {
		opt(category)
	}
	if err := category.validate(true); err != nil {
		return nil, NewValidationError(err.Error())
	}

	return category, nil
}

func NewCategoryID(id int64) NewCategoryOption {
	return func(c *CategoryModel) {
		c.id = id
	}
}

func NewCategoryName(name string) NewCategoryOption {
	return func(c *CategoryModel) {
		c.name = name
	}
}

func (category *CategoryModel) GetID() int64 {
	return category.id
}
func (category *CategoryModel) GetName() string {
	return category.name
}

func (category *CategoryModel) validate(shouldValidateID bool) *ValidationErrors {
	var errors []*ValidationError

	if shouldValidateID {
		if ve := category.isIDValid(); ve != nil {
			errors = append(errors, ve)
		}
	}

	if ve := category.isNameValid(); ve != nil {
		errors = append(errors, ve)
	}

	return validationErrorSliceToValidationErrors(errors)
}

func (category *CategoryModel) isIDValid() *ValidationError {
	if category.id <= 0 {
		return NewValidationError("category id must be greater than 0")
	}
	return nil
}
func (category *CategoryModel) isNameValid() *ValidationError {
	if category.name == "" {
		return NewValidationError("empty string in category name is not allowed")
	}
	if len(category.name) > categoryNameMaxLength {
		return NewValidationError(fmt.Sprintf("category name must be less than %d characters", categoryNameMaxLength))
	}
	return nil
}
