package post

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/model"
)

func Init(g *gin.Engine, postRepo model.PostRepository) {
	h := NewHandler(postRepo)

	RegisterRoutes(g.Group(""), h)
}
