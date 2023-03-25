package controller

import "github.com/google/uuid"

type GETAllUserResponse []getEachUser

type getEachUser struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}
