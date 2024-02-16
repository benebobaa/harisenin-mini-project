package middleware

import (
	"github.com/benebobaa/harisenin-mini-project/shared/util"
	"github.com/benebobaa/harisenin-mini-project/shared/util/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"net/http"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker, store *session.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Get the session
		sess, _ := store.Get(ctx)
		log.Println("sesion", sess)

		accessToken := sess.Get(AuthorizationHeaderKey)

		log.Println("acc token", accessToken)

		if ctx.Path() == "/" && accessToken == nil {
			return ctx.Next()
		}
		if accessToken != nil {

			payload, err := tokenMaker.VerifyToken(accessToken.(string))

			if err != nil {
				res, code := util.ConstructResponseError(http.StatusUnauthorized, err.Error())
				return ctx.Status(code).JSON(res)
			}

			payloadString := util.FromUsernameAndUUIDToString(payload.Username, payload.UUID)

			sess.Set(AuthorizationPayloadKey, payloadString)
			sess.Save()
			return ctx.Next()
		}

		return ctx.Redirect("/login")
	}
}
