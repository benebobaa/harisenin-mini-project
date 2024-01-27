package domain

import (
	"github.com/benebobaa/harisenin-mini-project/domain/repository"
	"github.com/benebobaa/harisenin-mini-project/domain/usecase"
	"github.com/benebobaa/harisenin-mini-project/infrastructure"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/go-playground/validator/v10"
)

type Domain struct {
	Validate     *validator.Validate
	QuoteUsecase usecase.QuoteUsecase
	TweetUsecase usecase.TweetUsecase
}

func ConstructDomain(c util.Config, validate *validator.Validate) Domain {
	databaseConn := infrastructure.NewDatabaseConnection(c.DBDsn)

	//Repository
	databaseRepository := repository.NewDatabaseRepository(databaseConn)
	tweetRepository := repository.NewTweetRepository(databaseConn)

	//Usecase
	quoteUsecase := usecase.NewQuoteUsecase(databaseRepository)
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository)

	return Domain{
		QuoteUsecase: &quoteUsecase,
		TweetUsecase: &tweetUsecase,
		Validate:     validate,
	}
}
