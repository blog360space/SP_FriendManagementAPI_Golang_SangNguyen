package friend

import "spapp/src/models/apimodels"

type GetFriendsInput struct {
	Email string `json:"email"`
}

type GetFriendsOutput struct {
	apimodels.ApiResult
	Count int `json:"count"`
	Friends []string `json:"friends"`
}