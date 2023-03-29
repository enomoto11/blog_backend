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

type post struct {
	ID         uuid.UUID
	CategoryID int64
	UserID     uuid.UUID
	Title      string
	Body       string
}

type createdPostResponse post

type allpostsResponse []post
