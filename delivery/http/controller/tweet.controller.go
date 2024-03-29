package controller

import (
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/dto/request"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/dto/response"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/middleware"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TweetController interface {
	FindAllTweet(ctx *fiber.Ctx) error
	CreateTweet(ctx *fiber.Ctx) error
	CommentTweet(ctx *fiber.Ctx) error
	FormTweet(ctx *fiber.Ctx) error
	CloseFormTweet(ctx *fiber.Ctx) error
}

type tweetControllerImpl struct {
	domain domain.Domain
}

func NewTweetController(domain domain.Domain) tweetControllerImpl {
	return tweetControllerImpl{domain: domain}
}

func (t *tweetControllerImpl) FindAllTweet(ctx *fiber.Ctx) error {

	var usernameUser string

	ses, err := t.domain.Store.Get(ctx)
	payload := ses.Get(middleware.AuthorizationPayloadKey)

	fmt.Println("PAYLOAD", payload)
	if payload != nil {
		usernameAndId, err := util.FromStringToUsernameAndUUID(payload.(string))

		if err != nil {
			resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch tweet")
			return ctx.Status(statusCode).JSON(resp)
		}

		usernameUser = usernameAndId.Username

	}

	tweets, err := t.domain.TweetUsecase.FindAllTweet()

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch tweet")
		return ctx.Status(statusCode).JSON(resp)
	}

	tweetResponses := response.NewTweetResponses(tweets)

	log.Printf("username %s", usernameUser)
	return ctx.Render("resource/views/home", fiber.Map{
		"Tweets":   tweetResponses,
		"Username": usernameUser,
	})
}

func (t *tweetControllerImpl) CreateTweet(ctx *fiber.Ctx) error {

	ses, _ := t.domain.Store.Get(ctx)
	payloadID := ses.Get(middleware.AuthorizationPayloadKey)

	var tweet request.TweetRequestDTO
	usernameAndId, err := util.FromStringToUsernameAndUUID(payloadID.(string))

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	tweet.UserID = usernameAndId.UserID

	if err := ctx.BodyParser(&tweet); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := t.domain.Validate.Struct(tweet); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	fileImage, err := ctx.FormFile("newImage")

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid file upload")
		return ctx.Status(statusCode).JSON(resp)
	}

	imageEntity, err := t.domain.AwsS3.UploadFile(fileImage)

	fmt.Println("ERROR UPLOAD", err)
	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid file upload")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := t.domain.TweetUsecase.CreateTweet(tweet.ToTweetEntity(imageEntity)); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save tweet")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

func (t *tweetControllerImpl) CommentTweet(ctx *fiber.Ctx) error {
	ses, _ := t.domain.Store.Get(ctx)
	payloadID := ses.Get(middleware.AuthorizationPayloadKey)
	fmt.Println("payload sess", payloadID)

	var comment request.CommentRequestDTO
	usernameAndId, err := util.FromStringToUsernameAndUUID(payloadID.(string))

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	comment.UserID = usernameAndId.UserID

	postId := ctx.Params("post_id")
	if err := ctx.BodyParser(&comment); err != nil {
		resp, _ := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Render("resource/views/home", fiber.Map{
			"Error": resp,
		})
	}

	comment.PostID = postId

	if err := t.domain.Validate.Struct(comment); err != nil {
		resp, _ := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Render("resource/views/home", fiber.Map{
			"Error": resp,
		})
	}

	if err := t.domain.TweetUsecase.CommentTweet(comment.ToCommentEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save comment")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

func (t *tweetControllerImpl) FormTweet(ctx *fiber.Ctx) error {
	ses, _ := t.domain.Store.Get(ctx)
	payloadID := ses.Get(middleware.AuthorizationPayloadKey)
	fmt.Println("payload sess", payloadID)

	return ctx.Render("resource/views/tweet_form", nil)
}

func (t *tweetControllerImpl) CloseFormTweet(ctx *fiber.Ctx) error {
	return ctx.Redirect("/")
}
