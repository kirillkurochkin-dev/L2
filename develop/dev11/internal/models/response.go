package models

type SuccessResponse struct {
	Result string `json:"result"`
}

type SuccessResponseAdd struct {
	Result []byte `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
