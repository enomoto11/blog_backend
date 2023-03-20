package model

import (
	"blog/api/utils"
	"errors"

	"strings"
)

type ValidationError struct {
	msg string
}

func (ve *ValidationError) Error() string {
	return ve.msg
}
func NewValidationError(msg string) *ValidationError {
	return &ValidationError{msg}
}

func IsValidationError(err error) bool {
	var validationErr *ValidationError
	return errors.As(err, &validationErr)
}

type ValidationErrors []*ValidationError

func (ves *ValidationErrors) Error() string {
	errMsgs := utils.MapSlice(*ves, func(v *ValidationError) string {
		if v == nil {
			return ""
		}
		return v.Error()
	})
	return strings.Join(errMsgs, "")
}

func validationErrorSliceToValidationErrors(errors []*ValidationError) *ValidationErrors {
	if len(errors) != 0 {
		err := ValidationErrors(errors)
		return &err
	}
	return nil
}
