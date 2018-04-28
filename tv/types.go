package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
)

type Show struct {
	helpers.Model
	ShowTitle 		string   `json:"show_title"`
	TvSeasons 		[]Season `json:"seasons,omitempty"`
}

type Season struct {
	helpers.Model
	SeasonNumber 	int         `json:"season_number"`
	ShowID 			uint        `json:"show_id"`
	Episodes 		[]Episode 	`json:"episodes,omitempty"`
}

// TODO: add a Show ID for convenience
type Episode struct {
	helpers.Model
	EpisodeTitle 	string			`json:"episode_title"`
	TvSeasonID 		uint			`json:"season_id"`
}

func InitTables(db *gorm.DB) error {
	return helpers.SafeCreateTable(db, &Show{}, &Season{}, &Episode{})
}
