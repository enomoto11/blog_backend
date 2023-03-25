package response

type AllUserResponse []user

type user struct {
	FirstName string
	LastName  string
	Email     string
}
