package router

import (
	"github.com/benebobaa/harisenin-mini-project/delivery/http/controller"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	quoteController := controller.NewQuoteController(domain)
	tweetController := controller.NewTweetController(domain)

	app.Get("/", quoteController.GetQuote)
	app.Post("/", quoteController.SaveQuote)

	app.Get("/tweets", tweetController.FindAllTweet)
	app.Get("/tweets/form", tweetController.FormTweet)
	app.Post("/tweets", tweetController.CreateTweet)
	app.Post("/tweets/comment/:post_id", tweetController.CommentTweet)

}
