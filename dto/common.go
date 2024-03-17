package dto

import "github.com/dgrijalva/jwt-go"

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
	ID          int    `json:"id"`
	Username    string `json:"username"`
	UserGroupID int    `json:"user_group_id"`
	Role        string `json:"role"`
	jwt.StandardClaims
}
