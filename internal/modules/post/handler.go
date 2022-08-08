package post

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/http"
	"go-template/internal/model"
)

type Handler struct {
	PostRepository model.PostRepository
}

func NewHandler(postRepo model.PostRepository) Handler {
	return Handler{PostRepository: postRepo}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreatePostRequest
	if err := c.Bind(&req); err != nil {
		http.HandleError(c, err, http.ValidationErr)
		return
	}

	post := model.Post{
		Title: req.Title,
		Body:  req.Body,
	}
	if err := h.PostRepository.Create(&post); err != nil {
		http.HandleError(c, err, http.InternalServerErr)
		return
	}

	http.Success(c, post, "create post successfully")
}

func (h *Handler) Update(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}

func (h *Handler) Get(c *gin.Context) {

}
