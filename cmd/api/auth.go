package main

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggles/pkg/user"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey = "id"
	port        string
)

// User demo

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func newAuthMiddleware(service user.UserService) *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(service),
		Authenticator:   authenticator(service),
		Authorizator:    authorizator(service),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func payloadFunc() func(data any) jwt.MapClaims {
	return func(data any) jwt.MapClaims {
		if v, ok := data.(*user.User); ok {
			return jwt.MapClaims{
				identityKey: v.UserName,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler(service user.UserService) func(c *gin.Context) any {
	return func(c *gin.Context) any {
		claims := jwt.ExtractClaims(c)

		username := claims[identityKey].(string)
		mongoUser, err := service.FindByUserName(username)

		if err != nil {
			return nil
		}

		return mongoUser
	}
}

func authenticator(service user.UserService) func(c *gin.Context) (any, error) {
	return func(c *gin.Context) (any, error) {
		apiKey := c.GetHeader("X-Api-Key")

		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		u, err := service.Authenticate(userID, password, apiKey)
		if err != nil {
			if err == user.ErrUserAuthenticationFailed {
				return nil, jwt.ErrFailedAuthentication
			}
			return nil, err
		}

		return u, nil
	}
}

func authorizator(service user.UserService) func(data any, c *gin.Context) bool {
	return func(data any, c *gin.Context) bool {
		routeName, ok := c.Get("routeName")
		if !ok {
			log.Printf("routeName not found for url %s\n", c.Request.URL.Path)
			return false
		}
		apiKey := c.GetHeader("X-Api-Key")
		u, ok := data.(*user.User)
		if ok && u.ApiKey == apiKey && service.HasPermission(u.UserName, routeName.(string)) {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}
