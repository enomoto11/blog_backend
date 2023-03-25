package request

import "github.com/google/uuid"

type POSTPostRequestBody struct {
	Title      string    `validate:"required"`
	Body       string    `validate:"required"`
	CategoryID int64     `validate:"required"`
	UserID     uuid.UUID `validate:"required"`
}
