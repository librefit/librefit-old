package database

import (
	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB holds the the database connection. I know, it's global variable. I
// think in this case might be justified ¯\_(ツ)_/¯
var DB *gorm.DB

// Connect creates a new gorm db connection.
func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./data/librefit.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db

	return db, nil
}

func Migrate() error {
	DB.AutoMigrate(
		&User{},
		&FoodDiary{},
		&FoodInventory{},
		&Measurement{},
		&Fluid{},
		&QuantityUnit{},
		&MealType{},
	)

	return nil
}

// GetDB use this function to get a connection
func GetDB() *gorm.DB {
	return DB
}

// SaveOne insert or update an element in the database, can be anything
func SaveOne(data interface{}) error {
	tx := DB.Save(data)
	if tx.Error != nil {
		spew.Dump(tx.Error)
		return tx.Error
	}

	return nil
}
