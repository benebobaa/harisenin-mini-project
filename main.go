package main

import (
	"embed"
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/delivery/http"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/benebobaa/harisenin-mini-project/shared/util/token"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/template/html/v2"
	httpLib "net/http"
)

//go:embed resource/*
var resourcefs embed.FS

func main() {
	config, err := util.LoadConfig(".")
	validate := validator.New()
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	store := session.New()

	domain := domain.ConstructDomain(config, validate, tokenMaker, store)
	engine := html.NewFileSystem(httpLib.FS(resourcefs), ".html")
	app := http.NewHttpDelivery(domain, engine, config)
	err = app.Listen(fmt.Sprintf(":%s", config.PortApp))
	if err != nil {
		return
	}
}
