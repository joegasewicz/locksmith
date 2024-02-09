package utilities

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func CreateDB() *gorm.DB {
	dsn := Config.PostgresConfig.GetPostgresConnStr()
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database failed to connect", err)
	}
	return database
}

var DB = CreateDB()
