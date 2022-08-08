package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-template/internal/config"
	"go-template/internal/db"
	"go-template/internal/modules/post"
	"go-template/internal/repository"
)

type Server struct{}

func (s *Server) Start() {
	g := gin.Default()
	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	cfg := config.LoadFromEnv()
	dbPG := db.DBPG{}
	dbPG.Connect(cfg.Postgres)

	postRepo := repository.NewPostRepository(dbPG.DB)
	post.Init(g, postRepo)

	err := g.Run(":8080")
	if err != nil {
		logrus.Fatal("Can't listen on port 8080")
		return
	}
}
