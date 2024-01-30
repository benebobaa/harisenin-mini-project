package usecase

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/benebobaa/harisenin-mini-project/domain/repository"
)

type TweetUsecase interface {
	FindAllTweet() ([]*entity.Post, error)
	CreateTweet(action entity.Post) error
	CommentTweet(action entity.Comment) error
	SaveImage(action entity.Image) error
}

type tweetUsecaseImpl struct {
	tweetRepository repository.TweetRepository
}

func NewTweetUsecase(tweetRepository repository.TweetRepository) tweetUsecaseImpl {
	return tweetUsecaseImpl{tweetRepository: tweetRepository}
}

func (t *tweetUsecaseImpl) FindAllTweet() ([]*entity.Post, error) {
	var tweets []*entity.Post
	if err := t.tweetRepository.FindAll(&tweets); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (t *tweetUsecaseImpl) CreateTweet(action entity.Post) error {

	if err := t.tweetRepository.Create(&action); err != nil {
		return err
	}

	return nil
}

func (t *tweetUsecaseImpl) CommentTweet(action entity.Comment) error {

	if err := t.tweetRepository.Comment(&action); err != nil {
		return err
	}

	return nil
}

func (t *tweetUsecaseImpl) SaveImage(action entity.Image) error {
	//TODO implement me
	panic("implement me")
}
