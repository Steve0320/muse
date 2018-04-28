package tv

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"log"
	"encoding/json"
)

// Served from /api/v1/tv
func HandleFullIndex(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		encoder := json.NewEncoder(w)
		var shows []TvShow

		err := db.Preload("TvSeasons").
			Preload("TvSeasons.TvEpisodes").
			Find(&shows).
			Error

		if err != nil {
			serverError(err, logger, w, encoder)
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
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(shows)

	})
}

func HandleSeasonsIndex(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		encoder := json.NewEncoder(w)
		var seasons []TvSeason

		err := db.Find(&seasons).Error
		if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(seasons)

	})
}

func HandleEpisodesIndex(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		encoder := json.NewEncoder(w)
		var seasons []TvEpisode

		err := db.Find(&seasons).Error
		if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(seasons)

	})
}

// Log and return an internal server error (code 500)
func serverError(err error, logger *log.Logger, w http.ResponseWriter, encoder *json.Encoder) {
	logger.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
	encoder.Encode("500: internal server error")
}