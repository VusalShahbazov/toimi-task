package repository

import (
	"context"
	"github.com/VusalShahbazov/toimi-api/internal/app/announcement"
	"github.com/VusalShahbazov/toimi-api/internal/entity"
	"github.com/go-pg/pg/v10"
)

type AnnRepo struct {
	DB *pg.DB
}

func NewAnnRepo(db *pg.DB) *AnnRepo {
	return &AnnRepo{DB: db}
}

//.WithContext(ctx)

func (r *AnnRepo) GetById(ctx context.Context, id int) (*entity.Announcement, error) {
	var ann entity.Announcement
	err := r.DB.Model(&ann).
		Where("id = ?", id).
		Relation("Photos").
		Select()
	if err != nil {
		return nil, err
	}
	return &ann, nil
}

func (r *AnnRepo) Get(ctx context.Context, opt announcement.AnnSearch) ([]entity.Announcement, error) {
	var ann []entity.Announcement

	err := r.DB.Model(&ann).
		Order(opt.Sort()).
		Limit(opt.GetLimit()).
		Offset(opt.GetOffset()).
		Relation("Photos").
		Select()
	if err != nil {
		return nil, err
	}

	return ann, nil
}

func (r *AnnRepo) Insert(ctx context.Context, ann entity.Announcement) (*entity.Announcement, error) {

	_, err := r.DB.Model(&ann).Insert(&ann)
	if err != nil {
		return nil, err
	}

	for i := range ann.Photos {
		ann.Photos[i].AnnouncementId = ann.Id
	}

	_, err = r.DB.Model(&ann.Photos).Insert(&ann.Photos)
	if err != nil {
		return nil, err
	}

	return &ann, nil
}
