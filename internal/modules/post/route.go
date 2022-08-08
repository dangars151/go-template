package post

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h Handler) {
	postGroup := r.Group("posts")
	postGroup.POST("", h.Create)
	postGroup.PUT("/:id", h.Update)
	postGroup.DELETE("/:id", h.Delete)
	postGroup.GET("", h.Get)
}
