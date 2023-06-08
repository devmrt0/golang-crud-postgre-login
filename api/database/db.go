package database

import (
	"cloudgobackend/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		db.LogMode(true)
		fmt.Printf("error connecting to the database:", err)
		fmt.Printf(config.DBURL)
		return nil, err
	}
	return db, nil
}
