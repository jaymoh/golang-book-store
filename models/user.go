package models

import (
	"hackinroms.com/books/models/commonfields"
)

type User struct {
	commonfields.Common
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
