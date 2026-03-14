package models

import "database/sql"

type UserModule struct {
	DB *sql.DB
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PaginatedResponse struct {
	Data       []User `json:"data"`
	TotalCount int    `json:"totalCount"`
	page       int    `json:"page"`
	pageSize   int    `json:"pageSize"`
}