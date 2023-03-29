package orm

import (
	"krittii/golang-api-jwt/orm/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func InitDB() {

	dsn := os.Getenv("POSTGRES_DNS")
	Db, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	Db.AutoMigrate(&model.User{})
}
