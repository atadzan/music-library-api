package repository

import (
	"context"

	"github.com/atadzan/music-library-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error)
}

type repository struct {
	db *pgxpool.Pool
}

func New(dbConn *pgxpool.Pool) Repository {
	return &repository{
		db: dbConn,
	}
}

func (r *repository) LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error) {
	return models.LibraryInfo{}, nil
}
