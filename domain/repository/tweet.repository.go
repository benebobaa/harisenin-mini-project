package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type TweetRepository interface {
	FindAll(data any, conditions ...any) error
	Create(value any) error
	Comment(value any) error
}

type tweetRepositoryImpl struct {
	database *gorm.DB
}

func NewTweetRepository(database *gorm.DB) tweetRepositoryImpl {
	return tweetRepositoryImpl{database: database}
}

func (t tweetRepositoryImpl) FindAll(data any, conditions ...any) error {
	//result := t.database.Find(data, conditions...)
	result := t.database.Preload("User").Preload("Comment.User").Order("created_at desc").Find(data, conditions...)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching tweet:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (t tweetRepositoryImpl) Create(value any) error {
	result := t.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating tweet:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (t tweetRepositoryImpl) Comment(value any) error {
	result := t.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating tweet:: %v", result.Error))
		return result.Error
	}

	return nil
}
