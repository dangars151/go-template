package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"go-template/internal/model"
)

type PostRepository struct {
	DB *pg.DB
}

func NewPostRepository(db *pg.DB) model.PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *model.Post) error {
	_, err := r.DB.Model(post).Insert()
	return err
}

func (r *PostRepository) CreateMany(posts []*model.Post) error {
	_, err := r.DB.Model(&posts).Insert()
	return err
}

func (r *PostRepository) Update(post *model.Post) error {
	_, err := r.DB.Model(post).WherePK().Update()
	return err
}

func (r *PostRepository) Delete(post *model.Post) error {
	_, err := r.DB.Model(post).WherePK().Delete()
	return err
}

func (r *PostRepository) FindByID(id uuid.UUID) (*model.Post, error) {
	return nil, nil
}

func (r *PostRepository) Find() ([]*model.Post, error) {
	return nil, nil
}
