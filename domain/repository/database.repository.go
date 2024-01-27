package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	FindRandomQuote(data any, conditions ...any) error
	Create(value any) error
}

type databaseRepositoryImpl struct {
	database *gorm.DB
}

func NewDatabaseRepository(database *gorm.DB) databaseRepositoryImpl {
	return databaseRepositoryImpl{
		database: database,
	}
}

func (this databaseRepositoryImpl) FindRandomQuote(data any, conditions ...any) error {
	result := this.database.Order("RANDOM()").Limit(1).Find(data, conditions...)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (this databaseRepositoryImpl) Create(value any) error {
	result := this.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating book:: %v", result.Error))
		return result.Error
	}

	return nil
}
