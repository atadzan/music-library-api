package services

import (
	"context"

	"github.com/atadzan/music-library-api/internal/models"
	"github.com/atadzan/music-library-api/internal/pkg/repository"
)

type Service interface {
	LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error)
	GetSongLyrics(ctx context.Context, param models.GetSongLyricsParams) (models.SongLyrics, error)
	AddSongToLibrary(ctx context.Context, param models.AddSongParams) (string, error)
	UpdateSong(ctx context.Context, params models.UpdateSongParams) (string, error)
	DeleteSong(ctx context.Context, songId int) (string, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error) {
	return s.repo.LibraryInfo(ctx, info)
}

func (s *service) GetSongLyrics(ctx context.Context, param models.GetSongLyricsParams) (models.SongLyrics, error) {
	return s.repo.GetSongLyrics(ctx, param)
}

func (s *service) AddSongToLibrary(ctx context.Context, param models.AddSongParams) (string, error) {
	return s.repo.AddSongToLibrary(ctx, param)
}

func (s *service) UpdateSong(ctx context.Context, params models.UpdateSongParams) (string, error) {
	return s.repo.UpdateSong(ctx, params)
}

func (s *service) DeleteSong(ctx context.Context, songId int) (string, error) {
	return s.repo.DeleteSong(ctx, songId)
}
