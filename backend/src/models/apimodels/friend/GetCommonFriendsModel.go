package friend

type GetCommonFriendsInput struct {
	Friends []string `json:"friends"`
}

type GetCommonFriendsOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
	Count int `json:"count"`
	Friends []string `json:"friends"`
}