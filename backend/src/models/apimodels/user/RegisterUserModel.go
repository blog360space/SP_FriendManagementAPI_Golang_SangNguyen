package user

import "spapp/src/models/apimodels"

type RegisterUserInput struct {
	Username string `json:"username"`
}

type RegisterUserOutput struct {
	apimodels.ApiResult
	Data struct {
		Id int `json:"id"`
		Username string `json:"username"`
	} `json:"data"`

}