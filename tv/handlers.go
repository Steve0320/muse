package tv

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"log"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
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
		var episodes []TvEpisode

		err := db.Find(&episodes).Error
		if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(episodes)

	})
}

func HandleShow(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		encoder := json.NewEncoder(w)

		// Get ID from URL
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			clientError(err, logger, w, encoder)
			return
		}

		var show TvShow

		err = db.Find(&show, id).Error
		if err != nil && err.Error() == "record not found" {
			notFoundError(err, logger, w, encoder)
			return
		} else if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(show)

	})
}

func HandleSeason(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		encoder := json.NewEncoder(w)

		// Get ID from URL
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			clientError(err, logger, w, encoder)
			return
		}

		var season TvSeason

		err = db.Find(&season, id).Error
		if err != nil && err.Error() == "record not found" {
			notFoundError(err, logger, w, encoder)
			return
		} else if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(season)

	})
}

func HandleEpisode(db *gorm.DB, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		encoder := json.NewEncoder(w)

		// Get ID from URL
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			clientError(err, logger, w, encoder)
			return
		}

		var episode TvEpisode

		err = db.Find(&episode, id).Error
		if err != nil && err.Error() == "record not found" {
			notFoundError(err, logger, w, encoder)
			return
		} else if err != nil {
			serverError(err, logger, w, encoder)
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(episode)

	})
}

// Log and return an internal server error (code 500)
func serverError(err error, logger *log.Logger, w http.ResponseWriter, encoder *json.Encoder) {
	logger.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
	encoder.Encode("500: internal server error")
}

// Log and return an client error (code 400)
func clientError(err error, logger *log.Logger, w http.ResponseWriter, encoder *json.Encoder) {
	logger.Print(err)
	w.WriteHeader(http.StatusBadRequest)
	encoder.Encode("400: malformed URL")
}

// Log and return a not found error (code 404)
func notFoundError(err error, logger *log.Logger, w http.ResponseWriter, encoder *json.Encoder) {
	logger.Print(err)
	w.WriteHeader(http.StatusNotFound)
	encoder.Encode("404: specified id does not exist")
}



