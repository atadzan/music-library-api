package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*

* Получение данных библиотеки с фильтрацией по всем полям и пагинацией
* Получение текста песни с пагинацией по куплетам
* Удаление песни
* Изменение данных песни
* Добавление новой песни в формате JSON

 */

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	app.RedirectTrailingSlash = true
	app.RedirectFixedPath = true

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	library := app.Group("/library")
	{
		library.GET("/info", getLibraryInfo)

		songs := library.Group("songs")
		{
			songs.GET("/:song_id/lyrics", getSongLyrics)
			songs.POST("/", addSongToLibrary)
			songs.DELETE("/:song_id", deleteSong)
			songs.PATCH("/:song_id", updateSong)
		}
	}
	return app
}
