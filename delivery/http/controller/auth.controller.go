package controller

import (
	"github.com/benebobaa/harisenin-mini-project/delivery/http/dto/request"
	"github.com/benebobaa/harisenin-mini-project/delivery/http/middleware"
	"github.com/benebobaa/harisenin-mini-project/domain"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/benebobaa/harisenin-mini-project/shared/util/token"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	SubmitLogin(ctx *fiber.Ctx) error
	SubmitRegister(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}

type authControllerImpl struct {
	domain domain.Domain
}

func NewAuthController(domain domain.Domain) authControllerImpl {
	return authControllerImpl{domain: domain}
}

func (a *authControllerImpl) Login(ctx *fiber.Ctx) error {
	return ctx.Render("resource/views/login", nil)
}

func (a *authControllerImpl) Register(ctx *fiber.Ctx) error {
	return ctx.Render("resource/views/register", nil)
}

func (a *authControllerImpl) SubmitLogin(ctx *fiber.Ctx) error {
	var login request.LoginRequestDTO
	if err := ctx.BodyParser(&login); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := a.domain.Validate.Struct(login); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	data, err := a.domain.AuthUsecase.Login(login.ToUserEntity())

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to login")
		return ctx.Status(statusCode).JSON(resp)
	}

	tokenPayload := token.UserPayload{Username: data.Username, UUID: data.ID.String()}
	accessToken, err := a.domain.TokenMaker.CreateToken(tokenPayload, time.Hour*24)

	if err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to login")
		return ctx.Status(statusCode).JSON(resp)
	}

	sess, _ := a.domain.Store.Get(ctx)
	sess.Set(middleware.AuthorizationHeaderKey, accessToken)
	sess.Save()
	return ctx.Redirect("/")
}

func (a *authControllerImpl) SubmitRegister(ctx *fiber.Ctx) error {
	var register request.RegisterRequestDTO
	if err := ctx.BodyParser(&register); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := a.domain.Validate.Struct(register); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := a.domain.AuthUsecase.Register(register.ToUserEntity()); err != nil {
		resp, statusCode := util.ConstructResponseError(fiber.StatusBadRequest, "Failed to register")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/login")
}

func (a *authControllerImpl) Logout(ctx *fiber.Ctx) error {

	sess, _ := a.domain.Store.Get(ctx)
	sess.Set(middleware.AuthorizationHeaderKey, nil)
	sess.Set(middleware.AuthorizationPayloadKey, nil)
	sess.Save()

	return ctx.Redirect("/")
}
