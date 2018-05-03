package tv

import (
	"muse2/helpers"
	"fmt"
	"github.com/manyminds/api2go"
)

type Season struct {
	helpers.Model
	SeasonNumber 	int         	`json:"season-number"`
	ShowID 			uint        	`json:"-"`
	Episodes 		[]Episode 		`json:"-"`
}

func (s Season) GetID() string {
	return fmt.Sprintf("%d", s.ID)
}

func (s Season) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	var season Season
	err := db.First(&season, ID).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: season}, nil
}

func (s Season) FindAll(req api2go.Request) (api2go.Responder, error) {
	var seasons []Season
	err := db.Find(&seasons).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: seasons}, nil
}
