package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct{}

func (s *Server) Start() {
	g := gin.Default()
	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	err := g.Run(":8080")
	if err != nil {
		logrus.Fatal("Can't listen on port 8080")
		return
	}
}