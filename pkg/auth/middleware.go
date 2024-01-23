package auth

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const CtxTokenKey = "userToken"
const CtxUserKey = "userId"

func MiddlewareWithSkipper(skippedPaths []string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(JwtSecret),
		TokenLookup: "cookie:" + CookieName,
		ContextKey:  CtxTokenKey,
		SuccessHandler: func(c echo.Context) {
			user := c.Get(CtxTokenKey).(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userId, _ := claims.GetIssuer()
			c.Set(CtxUserKey, userId)
		},
		Skipper: func(c echo.Context) bool {
			for _, path := range skippedPaths {
				if strings.HasSuffix(c.Path(), path) {
					return true
				}
			}
			return false
		},
	})
}

var Middleware echo.MiddlewareFunc = echojwt.WithConfig(echojwt.Config{
	SigningKey:  []byte(JwtSecret),
	TokenLookup: "cookie:" + CookieName,
	ContextKey:  CtxTokenKey,
	SuccessHandler: func(c echo.Context) {
		user := c.Get(CtxTokenKey).(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId, _ := claims.GetIssuer()
		c.Set(CtxUserKey, userId)
	},
	Skipper: func(c echo.Context) bool {

		return false
	},
})
