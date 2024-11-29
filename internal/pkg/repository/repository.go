package repository

import (
	"context"

	"github.com/atadzan/music-library-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	LibraryInfo(ctx context.Context, info models.LibraryInfoParams) (models.LibraryInfo, error)
	GetSongLyrics(ctx context.Context, param models.GetSongLyricsParams) (models.SongLyrics, error)
	AddSongToLibrary(ctx context.Context, param models.AddSongParams) (string, error)
	UpdateSong(ctx context.Context, params models.UpdateSongParams) (string, error)
	DeleteSong(ctx context.Context, songId int) (string, error)
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
func (r *repository) GetSongLyrics(ctx context.Context, param models.GetSongLyricsParams) (models.SongLyrics, error) {
	return models.SongLyrics{}, nil
}

func (r *repository) AddSongToLibrary(ctx context.Context, param models.AddSongParams) (string, error) {
	return "", nil
}

func (r *repository) UpdateSong(ctx context.Context, params models.UpdateSongParams) (string, error) {
	return "", nil
}

func (r *repository) DeleteSong(ctx context.Context, songId int) (string, error) {
	return "", nil
}
