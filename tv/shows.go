package tv

import (
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
	"muse2/helpers"
	"fmt"
	"github.com/manyminds/api2go/jsonapi"
)

type Show struct {
	helpers.Model
	db				*gorm.DB	`json:"-"`
	ShowTitle		string		`json:"show_title"`
	Seasons			[]Season	`json:"-"`
	Episodes		[]Episode	`json:"-"`
}

func (s Show) GetID() string {
	return fmt.Sprintf("%d", s.ID)
}

func NewShow(db *gorm.DB) *Show {
	return &Show{db: db}
}

func (s Show) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	var show Show
	err := s.db.First(&show, ID).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: show}, nil
}

func (s Show) FindAll(req api2go.Request) (api2go.Responder, error) {
	var shows []Show
	err := s.db.Find(&shows).Error
	if err != nil {
		return nil, err
	}
	return api2go.Response{Res: shows}, nil
}

func (s Show) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference {
		{
			Type: "seasons",
			Name: "seasons",
		},
	}
}

// TODO: fix
func (s Show) GetReferencedIds() []jsonapi.ReferenceID {

	result := []jsonapi.ReferenceID{}

	var ids []string
	err := s.db.Model(&Season{}).Where("show_id = ?", s.ID).Pluck("id", &ids).Error
	if err != nil {
		panic("oh no")
	}

	for _, id := range ids {
		result = append(result, jsonapi.ReferenceID{
			ID: id,
			Type: "seasons",
			Name: "seasons",
		})
	}

	return result

}
