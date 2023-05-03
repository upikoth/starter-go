package model

type ResponseErrorField struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ResponseSuccess struct {
	Success bool        `json:"success" default:"true"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Success bool                `json:"success" default:"false"`
	Error   *ResponseErrorField `json:"error"`
}
