package middleware

import (
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/benebobaa/harisenin-mini-project/shared/util/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker, store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if ctx.Path() == "/" {
			return ctx.Next()
		}

		// Get the session
		sess, _ := store.Get(ctx)

		accessToken := sess.Get(AuthorizationHeaderKey)
		if accessToken != nil {

			payload, err := tokenMaker.VerifyToken(accessToken.(string))

			if err != nil {
				res, code := util.ConstructResponseError(http.StatusUnauthorized, err.Error())
				return ctx.Status(code).JSON(res)
			}

			sess.Set(AuthorizationPayloadKey, fmt.Sprintf("%v", payload.UUID))
			sess.Save()
			return ctx.Next()
		}

		return ctx.Redirect("/login")
	}
}
