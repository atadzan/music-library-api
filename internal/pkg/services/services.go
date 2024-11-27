package services

import (
	"context"

	"github.com/atadzan/music-library-api/internal/models"
	"github.com/atadzan/music-library-api/internal/pkg/repository"
)

type Service interface {
	LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error)
}

type service struct {
	repo repository.Repository
}

func NewService() Service {
	return &service{}
}

func (s *service) LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error) {
	return models.LibraryInfo{}, nil
}
