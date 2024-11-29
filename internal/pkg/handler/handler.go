package handler

import (
	"net/http"

	"github.com/atadzan/music-library-api/internal/models"
	"github.com/atadzan/music-library-api/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

type Handler struct {
	service services.Service
}

func New(service services.Service) Handler {
	return Handler{
		service: service,
	}
}

// @Summary		Get library info
// @Tags			library
// @Description	Get library info
// @ID				get_library_info
// @Accept			json
// @Produce		json
// @Param		song_id	query	string	false	"song title"
// @Param		group	query	string	false	"song group"
// @Param		page	query	int		false	"page"
// @Param		limit	query	int		false	"limit"
// @Success		200			object	models.LibraryInfo
// @Failure		400			object	response	"<b>message:</b> invalid input params"
// @Failure		500			object	response	"<b>message:</b> internal server error"
// @Router			/library/info [get]
func (h *Handler) getLibraryInfo(c *gin.Context) {
	log.Debugf(c.Request.Context(), "getLibraryInfo request received")
	var input models.LibraryInfoParams
	if err := c.ShouldBind(&input); err != nil {
		log.Errorf(c.Request.Context(), "can't unmarshal input params. Err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "getLibraryInfo request get input params")
	resp, err := h.service.LibraryInfo(c.Request.Context(), input)
	if err != nil {
		log.Errorf(c.Request.Context(), "occurred error. Err: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, newMsgResponse("internal server error"))
		return
	}
	log.Infof(c.Request.Context(), "getLibraryInfo request successfully proceed")
	c.JSON(http.StatusOK, resp)
	return
}

// @Summary		Get song lyrics
// @Tags			library
// @Description	Get song lyrics
// @ID				get_song_lyrics
// @Accept			json
// @Produce		json
// @Param		song_id	query	string	false	"song title"
// @Param		group	query	string	false	"song group"
// @Param		page	query	int		false	"page"
// @Param		limit	query	int		false	"limit"
// @Success		200			object	models.LibraryInfo
// @Failure		400			object	response	"<b>message:</b> invalid input params"
// @Failure		500			object	response	"<b>message:</b> internal server error"
// @Router			/library/songs/song_id/lyrics [get]
func (h *Handler) getSongLyrics(c *gin.Context) {
	log.Debugf(c.Request.Context(), "getSongLyrics request received")
	var input models.GetSongLyricsParams
	if err := c.ShouldBind(&input); err != nil {
		log.Errorf(c.Request.Context(), "can't unmarshal input params. Err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "getSongLyrics request get input params")
	resp, err := h.service.GetSongLyrics(c.Request.Context(), input)
	if err != nil {
		log.Errorf(c.Request.Context(), "occurred error. Err: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "getSongLyrics request successfully proceed")
	c.JSON(http.StatusOK, resp)
	return
}

// @Summary		Add song to library
// @Tags			library
// @Description	Add song to library
// @ID				add_song_to_library
// @Accept			json
// @Produce		json
// @Param		input	body	models.AddSongParams	true	"song params"
// @Success		200			object	response 	"<b>message:</b> success"
// @Failure		400			object	response	"<b>message:</b> invalid input params"
// @Failure		500			object	response	"<b>message:</b> internal server error"
// @Router			/library/songs/song_id/lyrics [post]
func (h *Handler) addSongToLibrary(c *gin.Context) {
	log.Debugf(c.Request.Context(), "addSongToLibrary request received")
	var input models.AddSongParams
	if err := c.ShouldBind(&input); err != nil {
		log.Errorf(c.Request.Context(), "can't unmarshal input params. Err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "addSongToLibrary request get input params")
	resp, err := h.service.AddSongToLibrary(c.Request.Context(), input)
	if err != nil {
		log.Errorf(c.Request.Context(), "occurred error. Err: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "addSongToLibrary request successfully proceed")
	c.JSON(http.StatusOK, resp)
	return
}

// @Summary		Delete song
// @Tags			library
// @Description	Delete song
// @ID				delete_song_lyrics
// @Accept			json
// @Produce		json
// @Param		song_id	path	string	false	"song title"
// @Success		200			object	models.LibraryInfo
// @Failure		400			object	response	"<b>message:</b> invalid input params"
// @Failure		500			object	response	"<b>message:</b> internal server error"
// @Router			/library/songs/song_id [delete]
func (h *Handler) deleteSong(c *gin.Context) {
	log.Debugf(c.Request.Context(), "deleteSong request received")
	var input models.SongId
	if err := c.ShouldBind(&input); err != nil {
		log.Errorf(c.Request.Context(), "can't unmarshal input params. Err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "deleteSong request get input params")
	resp, err := h.service.DeleteSong(c.Request.Context(), input.SongId)
	if err != nil {
		log.Errorf(c.Request.Context(), "occurred error. Err: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "deleteSong request successfully proceed")
	c.JSON(http.StatusOK, resp)
	return
}

// @Summary		Update song
// @Tags			library
// @Description	Update song
// @ID				update_song_lyrics
// @Accept			json
// @Produce		json
// @Param		song_id	path	string	false	"song title"
// @Success		200			object	models.LibraryInfo
// @Failure		400			object	response	"<b>message:</b> invalid input params"
// @Failure		500			object	response	"<b>message:</b> internal server error"
// @Router			/library/songs/song_id [patch]
func (h *Handler) updateSong(c *gin.Context) {
	log.Debugf(c.Request.Context(), "updateSong request received")
	var input models.UpdateSongParams
	if err := c.ShouldBind(&input); err != nil {
		log.Errorf(c.Request.Context(), "can't unmarshal input params. Err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "updateSong request get input params")
	resp, err := h.service.UpdateSong(c.Request.Context(), input)
	if err != nil {
		log.Errorf(c.Request.Context(), "occurred error. Err: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, newMsgResponse("invalid input params"))
		return
	}
	log.Infof(c.Request.Context(), "updateSong request successfully proceed")
	c.JSON(http.StatusOK, resp)
	return
}
