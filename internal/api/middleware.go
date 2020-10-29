package api

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type User struct {
	Username string
	IsAdmin  bool
	UserID   uint
}

const identityKey = "id"

func newGinJWTMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     24 * time.Hour,
		MaxRefresh:  24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				s := jwt.MapClaims{
					"id":     v.Username,
					"role":   "user",
					"UserID": v.UserID,
				}
				if v.IsAdmin == true {
					s["role"] = "root"
				}
				return s
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := db.FindOneUser(&db.User{Username: loginVals.Username})
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			err = user.CheckPassword(loginVals.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &User{Username: user.Username, IsAdmin: user.IsAdmin, UserID: user.ID}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
