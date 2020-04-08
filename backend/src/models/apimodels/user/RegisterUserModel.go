package user

type RegisterUserInput struct {
	Username string `json:"username"`
}

type RegisterUserOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
	Data struct {
		Id int `json:"id"`
		Username string `json:"username"`
	}

}