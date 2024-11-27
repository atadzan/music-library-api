package models

type LibraryInfoParams struct {
	Song  string `form:"song" required:"true"`
	Group string `form:"group" required:"true"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}

type LibraryInfo struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type AddSongParams struct {
	Song        string `json:"song"`
	Group       string `json:"group"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type UpdateSongParams struct {
	Song        string  `path:"song"`
	Title       *string `json:"title"`
	Group       *string `json:"group"`
	ReleaseDate *string `json:"releaseDate"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}
