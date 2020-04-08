package friend


type BlockUserInput struct {
	Requestor string `json:"requestor"`
	Target string `json:"target"`
}

type BlockUserOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
}
