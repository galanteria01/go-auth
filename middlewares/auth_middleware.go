package middlewares

import (
	"example/go-auth/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Token"] != nil {
			bearerToken := c.Request.Header.Get("Token")
			_, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
					c.JSON(
						http.StatusUnauthorized,
						responses.Response{
							Status:  http.StatusUnauthorized,
							Message: "error",
							Data:    map[string]interface{}{"data": "You are unauthorized"},
						},
					)
				}
				return "", nil
			})

			if err != nil {
				c.JSON(
					http.StatusUnauthorized,
					responses.Response{
						Status:  http.StatusUnauthorized,
						Message: "error",
						Data:    map[string]interface{}{"data": "You are unauthorized"},
					},
				)
				return
			}

			c.Next()

		} else {
			c.JSON(
				http.StatusUnauthorized,
				responses.Response{
					Status:  http.StatusUnauthorized,
					Message: "error",
					Data:    map[string]interface{}{"data": "You are unauthorized"},
				},
			)
		}
	}
}
