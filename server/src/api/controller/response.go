package controller

import "github.com/google/uuid"

type GETAllUserResponse []getEachUser

type getEachUser struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}

type createdUserResponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}

type createdCategoryResponse struct {
	ID   int64
	Name string
}

type createdPostResponse struct {
	ID         uuid.UUID
	CategoryID int64
	UserID     uuid.UUID
	Title      string
	Body       string
}
