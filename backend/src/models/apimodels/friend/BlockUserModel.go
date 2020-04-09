package friend

import "spapp/src/models/apimodels"

type BlockUserInput struct {
	Requestor string `json:"requestor"`
	Target string `json:"target"`
}

type BlockUserOutput struct {
	apimodels.ApiResult
}