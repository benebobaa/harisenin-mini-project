package infrastructure

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabaseConnection(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Quote{}, &entity.Post{}, &entity.Comment{}); err != nil {
		log.Fatal(err)
	}

	return db
}
