package model

import "github.com/google/uuid"

type Post struct {
	tableName struct{} `pg:"posts,alias:posts"`
	Base
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostRepository interface {
	Create(post *Post) error
	Update(post *Post) error
	Delete(post *Post) error
	FindByID(id uuid.UUID) (*Post, error)
	Find() ([]*Post, error)
}
