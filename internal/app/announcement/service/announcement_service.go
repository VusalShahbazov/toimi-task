package service

import (
	"context"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement/repository"
	"github.com/VusalShahbazov/toimi-api/internal/entity"
)

type AnnSrv struct {
	annRepo repository.AnnRepo
}

func NewAnnSrv(annRepo repository.AnnRepo) *AnnSrv {
	return &AnnSrv{annRepo: annRepo}
}

func (r *AnnSrv) GetById(ctx context.Context, id int) (*entity.Announcement, error) {
	return r.annRepo.GetById(ctx, id)
}

func (r *AnnSrv) Get(ctx context.Context, opt announcement.AnnSearch) ([]entity.Announcement, error) {
	return r.annRepo.Get(ctx, opt)
}

func (r *AnnSrv) Insert(ctx context.Context, ann entity.Announcement) (*entity.Announcement, error) {
	return r.annRepo.Insert(ctx, ann)
}
