package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type ApiServer struct {
	Config *Config
	Router *gin.Engine
	DB     *pg.DB
}

func NewApiServer(config *Config) *ApiServer {
	return &ApiServer{
		Config: config,
		Router: gin.Default(),
	}
}

func (s *ApiServer) Run() error {
	// Init database
	err := s.InitDatabase()
	if err != nil {
		return err
	}

	err = s.initHandlers()
	if err != nil {
		return err
	}

	err = s.Router.Run(s.Config.BindAddr)

	return err
}
