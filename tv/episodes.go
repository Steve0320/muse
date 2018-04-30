package tv

import (
	"muse2/helpers"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

type Episode struct {
	helpers.Model
	db				*gorm.DB		`json:"-"`
	EpisodeTitle 	string			`json:"episode_title"`
	SeasonID 		uint			`json:"-"`
	ShowID			uint			`json:"-"`
}

func (e Episode) GetID() string {
	return fmt.Sprintf("%d", e.ID)
}

func NewEpisode(db *gorm.DB) *Episode {
	return &Episode{db: db}
}

func (e Episode) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	var epsiode Episode
	err := e.db.First(&epsiode, ID).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: epsiode}, nil
}

func (e Episode) FindAll(req api2go.Request) (api2go.Responder, error) {
	var episodes []Episode
	err := e.db.Find(&episodes).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: episodes}, nil
}
