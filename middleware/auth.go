package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if errEnv := godotenv.Load(); errEnv != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			return
		}

		secret := os.Getenv("SECRET")
		author := ctx.Request.Header.Get("Auhtorization")
		tokenString := strings.Replace(author, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorize"})
			return
		}

		Token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != t.Method {
				return nil, errors.New("Unauthorize")
			}
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorize"})
			return
		}

		claims, ok := Token.Claims.(jwt.MapClaims)

		if !Token.Valid || !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": http.StatusText(http.StatusForbidden)})
			return
		}

		exp := claims["exp"].(int64)

		if time.Now().After(time.Unix(exp, 0)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": "Token Is Expired"})
			return
		}

		id := claims["Uid"].(string)

		ctx.Set("Id", id)
		ctx.Next()
	}
}
