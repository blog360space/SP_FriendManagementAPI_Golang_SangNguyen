package user

type RegisterUserInput struct {
	Username string
}

type RegisterUserOutput struct {
	Id int32
	Username string
}