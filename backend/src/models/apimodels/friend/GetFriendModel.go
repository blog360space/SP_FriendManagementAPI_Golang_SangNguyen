package friend

type GetFriendsInput struct {
	Email string `json:"email"`
}

type GetFriendsOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
	Count int `json:"count"`
	Friends []string `json:"friends"`
}