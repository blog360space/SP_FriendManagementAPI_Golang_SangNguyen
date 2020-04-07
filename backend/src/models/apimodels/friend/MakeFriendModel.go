package friend

type MakeFriendInput struct {
	Friends []string `json:"friends"`
}

type MakeFriendOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
}