package helpers

import (
	"github.com/jinzhu/gorm"
	"time"
)

// More flexible version of gorm.Model
type Model struct {
	ID uint 				`gorm:"primary_key" json:"-"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"-"`
}

// Create tables if they don't already exist
func MassCreateTable(db *gorm.DB, types ...interface{}) error {

	for _, t := range types {
		err := db.AutoMigrate(t).Error
		if err != nil {
			return err
		}
	}

	return nil

}
