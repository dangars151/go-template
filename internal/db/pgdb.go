package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go-template/internal/config"
)

type DBPG struct {
	DB *pg.DB
}

func (s *DBPG) Connect(pc *config.PostgresConfig) {
	connectionString := pc.GetPGConnectionString()
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		logrus.Panic(fmt.Errorf("parse go-pg connection string err: %w", err))
	}

	s.DB = pg.Connect(opt)
	if err = s.DB.Ping(context.Background()); err != nil {
		logrus.Panic(fmt.Errorf("connect to postgres err: %w", err))
	}

	logrus.Info("connect db pg successfully")
}
