package friend

import "spapp/src/models/apimodels"

type GetCommonFriendsInput struct {
	Friends []string `json:"friends"`
}

type GetCommonFriendsOutput struct {
	apimodels.ApiResult
	Count int `json:"count"`
	Friends []string `json:"friends"`
}