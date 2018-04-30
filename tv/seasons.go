package tv

import (
	"muse2/helpers"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

type Season struct {
	helpers.Model
	db				*gorm.DB		`json:"-"`
	SeasonNumber 	int         	`json:"season_number"`
	ShowID 			uint        	`json:"-"`
	Episodes 		[]Episode 		`json:"-"`
}

func (s Season) GetID() string {
	return fmt.Sprintf("%d", s.ID)
}

func NewSeason(db *gorm.DB) *Season {
	return &Season{db: db}
}

func (s Season) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	var season Season
	err := s.db.First(&season, ID).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: season}, nil
}

func (s Season) FindAll(req api2go.Request) (api2go.Responder, error) {
	var seasons []Season
	err := s.db.Find(&seasons).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: seasons}, nil
}
