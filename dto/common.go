package dto

import "github.com/golang-jwt/jwt/v5"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

type ResponseList struct {
	Response
	Pagination *Pagination `json:"pagination"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.Claims
}
