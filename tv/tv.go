package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
)

var db *gorm.DB

// Set the global database and create tables
func InitDB(database *gorm.DB) error {
	db = database
	return helpers.MassCreateTable(database, &Show{}, &Season{}, &Episode{})
}
