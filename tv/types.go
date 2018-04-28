package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
)

type Show struct {
	helpers.Model
	ShowTitle 		string   		`json:"show_title"`
	Seasons 		[]Season 		`json:"seasons,omitempty"`
	Episodes		[]Episode		`json:"episodes,omitempty"`
}

type Season struct {
	helpers.Model
	SeasonNumber 	int         	`json:"season_number"`
	ShowID 			uint        	`json:"show_id"`
	Episodes 		[]Episode 		`json:"episodes,omitempty"`
}

type Episode struct {
	helpers.Model
	EpisodeTitle 	string			`json:"episode_title"`
	SeasonID 		uint			`json:"season_id"`
	ShowID			uint			`json:"show_id"`
}

func InitTables(db *gorm.DB) error {
	return helpers.SafeCreateTable(db, &Show{}, &Season{}, &Episode{})
}
