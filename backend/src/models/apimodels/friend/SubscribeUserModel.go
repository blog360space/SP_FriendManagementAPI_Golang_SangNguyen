package friend

type SubscribeUserInput struct {
	Requestor string `json:"requestor"`
	Target string `json:"target"`
}

type SubscribeUserOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
}