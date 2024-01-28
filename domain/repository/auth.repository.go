package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type AuthRepository interface {
	CreateUser(value any) error
	FindByUsername(value any, username string) error
}

type authRepositoryImpl struct {
	database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) authRepositoryImpl {
	return authRepositoryImpl{database: database}
}

func (a authRepositoryImpl) CreateUser(value any) error {
	result := a.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating user regist:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (a authRepositoryImpl) FindByUsername(value any, username string) error {
	result := a.database.Where("username = ?", username).Find(&value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error find user:: %v", result.Error))
		return result.Error
	}

	return nil
}
