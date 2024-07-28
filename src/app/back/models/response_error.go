package models

type ResponseError struct {
	HTTPStatusCode int    `json:"-"`
	ErrorCode      string `json:"error_code"`
	Message        string `json:"message"`
}
