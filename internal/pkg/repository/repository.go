package repository

import (
	"context"

	"github.com/atadzan/music-library-api/internal/models"
)

type Repository interface {
	LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error)
}

type repository struct {
}

func New() Repository {
	return &repository{}
}

func (r *repository) LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error) {
	return models.LibraryInfo{}, nil
}
