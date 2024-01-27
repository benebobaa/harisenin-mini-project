package http

import (
	"github.com/benebobaa/harisenin-mini-project/delivery/http/router"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func NewHttpDelivery(domain domain.Domain, engine *html.Engine, c util.Config) *fiber.App {
	config := fiber.Config{
		AppName:           c.AppName,
		EnablePrintRoutes: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		Views:             engine,
	}

	if c.GOEnv == "production" {
		config.Prefork = true
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	router.NewRouter(app, domain)

	return app
}
