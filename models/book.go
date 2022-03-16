package models

import (
	"hackinroms.com/books/models/commonfields"
)

type Book struct {
	commonfields.Common
	Title     string `json:"title"`
	Author    string `json:"author"`
	PageCount int    `json:"page_count"`
}
