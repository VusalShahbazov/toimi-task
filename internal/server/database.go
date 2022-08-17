package server

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func (s *ApiServer) InitDatabase() error {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%v:%v", s.Config.DBConfig.Host, s.Config.DBConfig.Port),
		User:     s.Config.DBConfig.Username,
		Password: s.Config.DBConfig.Password,
		Database: s.Config.DBConfig.Database,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return err
	}

	s.DB = db

	return nil
}
