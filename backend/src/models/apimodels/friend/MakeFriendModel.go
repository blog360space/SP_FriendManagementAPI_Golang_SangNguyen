package friend

import "spapp/src/models/apimodels"

type MakeFriendInput struct {
	Friends []string `json:"friends"`
}

type MakeFriendOutput struct {
	apimodels.ApiResult
}