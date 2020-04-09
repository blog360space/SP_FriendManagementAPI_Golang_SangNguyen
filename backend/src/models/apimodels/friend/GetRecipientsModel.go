package friend

import "spapp/src/models/apimodels"

type GetRecipientsInput struct {
	Sender string `json:"sender"`
	Text string `json:"text"`
}

type GetRecipientsOutput struct {
	apimodels.ApiResult
	Recipients []string `json:"recipients"`
}