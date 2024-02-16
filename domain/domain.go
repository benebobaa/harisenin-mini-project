package domain

import (
	"github.com/benebobaa/harisenin-mini-project/domain/repository"
	"github.com/benebobaa/harisenin-mini-project/domain/usecase"
	"github.com/benebobaa/harisenin-mini-project/infrastructure"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/benebobaa/harisenin-mini-project/shared/util/token"
	"github.com/benebobaa/harisenin-mini-project/shared/util/upload_image"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Domain struct {
	Validate     *validator.Validate
	TokenMaker   token.Maker
	Store        *session.Store
	AwsS3        upload_image.AwsS3Action
	TweetUsecase usecase.TweetUsecase
	AuthUsecase  usecase.AuthUsecase
}

func ConstructDomain(c util.Config, validate *validator.Validate, tokenMaker token.Maker, store *session.Store, awsS3 upload_image.AwsS3Action) Domain {
	databaseConn := infrastructure.NewDatabaseConnection(c.DBDsn)

	//Repository
	tweetRepository := repository.NewTweetRepository(databaseConn)
	authRepository := repository.NewAuthRepository(databaseConn)

	//Usecase
	tweetUsecase := usecase.NewTweetUsecase(tweetRepository)
	authUsecase := usecase.NewAuthUsecase(authRepository)

	return Domain{
		Validate:     validate,
		TokenMaker:   tokenMaker,
		Store:        store,
		AwsS3:        awsS3,
		TweetUsecase: &tweetUsecase,
		AuthUsecase:  &authUsecase,
	}
}
