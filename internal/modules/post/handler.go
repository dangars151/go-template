package post

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/http"
	"go-template/internal/model"
	"go-template/utils/postutil"
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
		http.HandleError(c, err, http.PostgresErr)
		return
	}

	http.Success(c, post, "create post successfully")
}

func (h *Handler) CreateFromApi(c *gin.Context) {
	postsApi, err := postutil.GetPosts()
	if err != nil {
		http.HandleError(c, err, http.InternalServerErr)
		return
	}

	posts := make([]*model.Post, 0)
	for _, p := range postsApi {
		posts = append(posts, &model.Post{
			Title: p.Title,
			Body:  p.Body,
		})
	}

	if err = h.PostRepository.CreateMany(posts); err != nil {
		http.HandleError(c, err, http.PostgresErr)
		return
	}

	http.Success(c, nil, "Import from api successfully")
}

func (h *Handler) Update(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}

func (h *Handler) Get(c *gin.Context) {

}
