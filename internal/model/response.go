package model

type SuccessResponse struct {
	Redirect string `json:"redirect"`
	Noko     bool   `json:"noko"`
	ID       string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
