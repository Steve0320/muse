package tv

import (
	"github.com/manyminds/api2go"
	"muse2/helpers"
	"fmt"
	"net/http"
	"strconv"
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"log"
)

type Show struct {
	helpers.Model
	Title			string		`json:"title"`
	TvID			string		`json:"tvid"`
	Description		string		`json:"description"`
	CoverURL		string		`json:"cover-url"`
	Seasons			[]Season	`json:"-"`
	Episodes		[]Episode	`json:"-"`
}

//################################ Implement MarshalIdentifier ################################//

func (s Show) GetID() string {
	return fmt.Sprintf("%d", s.ID)
}

//################################ Implement UnmarshalIdentifier ################################//

func (s *Show) SetID(id string) error {

	// Let database handle ID
	if id == "" {
		return nil
	}

	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}

	s.ID = uint(u64)
	return nil

}

//################################ Implement FindAll ################################//

func (s Show) FindAll(req api2go.Request) (api2go.Responder, error) {
	var shows []Show
	err := db.Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: shows}, nil
}

//################################ Implement CRUD ################################//

// Create
func (s Show) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {

	show, ok := obj.(Show)
	if !ok {
		err := errors.New("invalid instance given")
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}

	err := db.Create(&show).Error
	if err != nil {
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	return &api2go.Response{Res: show, Code: http.StatusCreated}, nil

}

// Read
func (s Show) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {

	var show Show

	err := db.First(&show, ID).Error
	if err != nil {
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return api2go.Response{Res: show}, nil

}

// Update
func (s Show) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {

	show, ok := obj.(Show)
	if !ok {
		err := errors.New("invalid instance given")
		return &api2go.Response{}, api2go.NewHTTPError(err, "invalid instance given", http.StatusBadRequest)
	}

	err := db.Save(&show).Error
	if err != nil {
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}

	return &api2go.Response{Res: show, Code: http.StatusNoContent}, err

}

// Delete
func (s Show) Delete(id string, r api2go.Request) (api2go.Responder, error) {

	err := db.Delete(&Show{}, id).Error
	return &api2go.Response{Code: http.StatusNoContent}, err

}

func (s Show) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference {
		{
			Type: "seasons",
			Name: "seasons",
		},
		//{
		//	Type: "episodes",
		//	Name: "episodes",
		//},
	}
}

func (s Show) GetReferencedIDs() []jsonapi.ReferenceID {

	var result []jsonapi.ReferenceID
	var seasonIDs []string
	//var episodeIDs []string
	var err error

	// Fetch associated season IDs
	err = db.Model(&Season{}).Where("show_id = ?", s.ID).Pluck("id", &seasonIDs).Error
	if err != nil {
		log.Print("ERROR: get referenced season IDs for Show failed")
		return result
	}

	// Fetch associated episode IDs
	//err = db.Model(&Episode{}).Where("show_id = ?", s.ID).Pluck("id", &episodeIDs).Error
	//if err != nil {
	//	log.Print("ERROR: get referenced episode IDs for Show failed")
	//	return result
	//}

	// Add season IDs
	for _, id := range seasonIDs {
		result = append(result, jsonapi.ReferenceID{
			ID: id,
			Type: "seasons",
			Name: "seasons",
		})
	}

	// Add episode IDs
	//for _, id := range episodeIDs {
	//	result = append(result, jsonapi.ReferenceID{
	//		ID: id,
	//		Type: "episodes",
	//		Name: "episodes",
	//	})
	//}

	return result

}

/*
func (s Show) GetReferencedStructs() []jsonapi.MarshalIdentifier {

	var results []jsonapi.MarshalIdentifier
	var seasons []Season
	//var episodes []Episode

	err := db.Find(&seasons).Where("show_id = ?", s.ID).Error
	if err != nil {
		log.Print("ERROR: get referenced seasons for Show failed")
	}

	//err = db.Find(&episodes).Where("show_id = ?", s.ID).Error
	//if err != nil {
	//	log.Print("ERROR: get referenced episodes for Show failed")
	//}

	for _, season := range seasons {
		results = append(results, season)
	}

	//for _, episode := range episodes {
	//	results = append(results, episode)
	//}

	return results

}
*/