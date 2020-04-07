package user

type RegisterUserInput struct {
	Username string `json:"username"`
}

type RegisterUserOutput struct {
	Id int `json:"id"`
	Username string `json:"username"`
}