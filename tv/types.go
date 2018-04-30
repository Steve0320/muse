package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
	"github.com/manyminds/api2go"
	"fmt"
)

type Show struct {
	helpers.Model
	ShowTitle 		string   		`json:"show_title"`
	//Seasons 		[]Season 		`json:"seasons,omitempty"`
	//Episodes		[]Episode		`json:"episodes,omitempty"`
}

func (s Show) GetID() string  {
	return fmt.Sprintf("%d", s.ID)
}

type ShowSource struct {
	db *gorm.DB
}

func NewShowSource(db *gorm.DB) *ShowSource {
	return &ShowSource{db}
}

// FindOne returns an object by its ID
// Possible Responder success status code 200
func (s *ShowSource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	var show Show
	err := s.db.Find(&show, ID).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: show}, nil
}

/*
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
*/

func InitTables(db *gorm.DB) error {
	//return helpers.SafeCreateTable(db, &Show{}, &Season{}, &Episode{})
	return helpers.SafeCreateTable(db, &Show{})
}
