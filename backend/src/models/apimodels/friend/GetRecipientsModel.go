package friend

type GetRecipientsInput struct {
	Sender string `json:"sender"`
	Text string `json:"text"`
}

type GetRecipientsOutput struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
	Recipients []string `json:"recipients"`
}