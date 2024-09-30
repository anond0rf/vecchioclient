package model

type ReplyResponse struct {
	Redirect string `json:"redirect"`
	Noko     bool   `json:"noko"`
	ID       string `json:"id"`
}

type ReplyErrorResponse struct {
	Error string `json:"error"`
}
