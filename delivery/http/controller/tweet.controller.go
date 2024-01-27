package controller

import (
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/dto/request"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/dto/response"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/gofiber/fiber/v2"
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
	tweets, err := t.domain.TweetUsecase.FindAllTweet()

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to fetch tweet")
		return ctx.Status(statusCode).JSON(resp)
	}

	tweetResponse := response.NewTweetResponse(tweets)

	fmt.Println(tweetResponse)
	return ctx.Render("resource/views/tweet", fiber.Map{
		"Tweets": tweetResponse,
	})
}

func (t *tweetControllerImpl) CreateTweet(ctx *fiber.Ctx) error {
	var tweet request.TweetRequestDTO
	if err := ctx.BodyParser(&tweet); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := t.domain.Validate.Struct(tweet); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := t.domain.TweetUsecase.CreateTweet(tweet.ToTweetEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save tweet")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/tweets")
}

func (t *tweetControllerImpl) CommentTweet(ctx *fiber.Ctx) error {
	var comment request.CommentRequestDTO
	postId := ctx.Params("post_id")
	if err := ctx.BodyParser(&comment); err != nil {
		resp, _ := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Render("resource/views/tweet", fiber.Map{
			"Error": resp,
		})
	}

	comment.PostID = postId

	if err := t.domain.Validate.Struct(comment); err != nil {
		resp, _ := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Render("resource/views/tweet", fiber.Map{
			"Error": resp,
		})
	}

	if err := t.domain.TweetUsecase.CommentTweet(comment.ToCommentEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to save comment")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/tweets")
}

func (t *tweetControllerImpl) FormTweet(ctx *fiber.Ctx) error {
	return ctx.Render("resource/views/tweet_form", nil)
}

func (t *tweetControllerImpl) CloseFormTweet(ctx *fiber.Ctx) error {
	return ctx.Redirect("/tweets")
}
