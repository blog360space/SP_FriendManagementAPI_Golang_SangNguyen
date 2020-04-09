package apimodels

type ApiResult struct {
	Success bool `json:"success"`
	Msgs []string `json:"msgs"`
}