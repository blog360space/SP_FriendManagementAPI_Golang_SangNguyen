package user

type RegisterUserInput struct {
	Username string `json:"username"`
}

type RegisterUserOutput struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
}