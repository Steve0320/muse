package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
)

type TvShow struct {
	helpers.Model
	ShowTitle 		string 		`json:"show_title"`
	TvSeasons 		[]TvSeason 	`json:"seasons,omitempty"`
}

type TvSeason struct {
	helpers.Model
	SeasonNumber 	int				`json:"season_number"`
	TvShowID 		uint			`json:"tv_show_id"`
	TvEpisodes 		[]TvEpisode		`json:"episodes,omitempty"`
}

type TvEpisode struct {
	helpers.Model
	EpisodeTitle 	string			`json:"episode_title"`
	TvSeasonID 		uint			`json:"tv_season_id"`
}

func InitTables(db *gorm.DB) error {
	return helpers.SafeCreateTable(db, &TvShow{}, &TvSeason{}, &TvEpisode{})
}
