package friend

import "spapp/src/models/apimodels"

type SubscribeUserInput struct {
	Requestor string `json:"requestor"`
	Target string `json:"target"`
}

type SubscribeUserOutput struct {
	apimodels.ApiResult
}