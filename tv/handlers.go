package tv

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"log"
	"encoding/json"
)

// Served from /api/v1/tv
func HandleIndex(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		encoder := json.NewEncoder(w)
		var shows []TvShow

		err := db.Preload("TvSeasons").
			Preload("TvSeasons.TvEpisodes").
			Find(&shows).
			Error

		if err != nil {
			logger.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode("500: internal server error")
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(shows)

	})
}

// Served from /api/v1/tv/shows
func HandleShowsIndex(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		encoder := json.NewEncoder(w)
		var shows []TvShow

		err := db.Find(&shows).Error
		if err != nil {
			logger.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode("500: internal server error")
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(shows)

	})
}
