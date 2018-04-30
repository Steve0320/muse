package tv

import (
	"github.com/jinzhu/gorm"
	"muse2/helpers"
)

func InitTables(db *gorm.DB) error {
	return helpers.MassCreateTable(db, &Show{}, &Season{}, &Episode{})
}
