package jwt

import (
	"net/http"
	"um_sys/pkg/e"
	"um_sys/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {

		var code int
		var data interface{}

		code = e.SUCCESS

		token := context.Query("token")

		if token == "" {
			code = e.INVALID_PARAM
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			context.Abort()
			return
		}

		context.Next()
	}
}
