package server

import (
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement/handler"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement/repository"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement/service"
)

func (s *ApiServer) initHandlers() error {

	annRepo := repository.NewAnnRepo(s.DB)
	annSrv := service.NewAnnSrv(*annRepo)
	handler.InitAnnouncementHandler(s.Router, *annSrv)
	return nil
}
