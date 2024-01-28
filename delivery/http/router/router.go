package router

import (
	"github.com/benebobaa/harisenin-mini-project/delivery/http/controller"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/middleware"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, domain domain.Domain) {

	tweetController := controller.NewTweetController(domain)
	authController := controller.NewAuthController(domain)

	//Exclude From Middleware
	app.Get("/login", authController.Login)
	app.Get("/register", authController.Register)
	app.Post("/login", authController.SubmitLogin)
	app.Post("/register", authController.SubmitRegister)

	//Group Middleware
	home := app.Group("", middleware.AuthMiddleware(domain.TokenMaker, domain.Store))

	home.Get("/", tweetController.FindAllTweet)
	home.Get("/form", tweetController.FormTweet)
	home.Post("/form", tweetController.CreateTweet)
	home.Post("/:post_id", tweetController.CommentTweet)

}
