package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggles/internal/envs"
	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
	"github.com/markelca/toggles/pkg/user"
)

var startTime time.Time // Used to calculate uptime
var users map[string]user.User
var params *envs.ConnectionParams
var userRepo user.UserRepository

func init() {
	var paramErr []error
	params, paramErr = envs.GetConnectionParams(envs.EnvNames{
		AppPort:   "APP_PORT",
		RedisHost: "REDIS_HOST",
		RedisPort: "REDIS_PORT",
		MongoHost: "MONGO_HOST",
		MongoPort: "MONGO_PORT",
	})
	if len(paramErr) > 0 {
		envs.PrintFatalErrors(paramErr)
	}

	// repository := storage.NewMemoryRepository()
	startTime = time.Now()
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var err error
	userRepo, err = user.NewUserMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

}
func uptime() time.Duration {
	return time.Since(startTime)
}

func main() {
	engine := gin.Default()
	// the jwt middleware
	authMiddleware, err := jwt.New(newAuthMiddleware(userRepo))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// register middleware
	engine.Use(handlerMiddleWare(authMiddleware))

	// register route
	registerRoute(engine, authMiddleware)

	// start http server
	if err = http.ListenAndServeTLS(":"+port, "./testdata/selfsigned.crt", "./testdata/selfsigned.key", engine); err != nil {
		log.Fatal(err)
	}
}

func registerRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	db, err := flags.NewFlagMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

	authMiddlewareFunc := authMiddleware.MiddlewareFunc()

	repository := storage.NewRedisClient(params.RedisHost, params.RedisPort)
	service := flags.NewFlagService(repository, db)
	controller := NewFlagController(service)

	r.NoRoute(handleNoRoute())
	r.GET("/health-check", RouteName("health-check"), healthHandler)
	r.POST("/login", authMiddleware.LoginHandler, RouteName("login"))
	r.GET("/refresh_token", RouteName("refresh_token"), authMiddlewareFunc, authMiddleware.RefreshHandler)

	r.GET("/me", RouteName("get_me"), authMiddlewareFunc, meHandler)

	r.GET("/flags", RouteName("get_flags"), authMiddlewareFunc, controller.ListFlags)
	r.GET("/flags/:flagid", RouteName("get_flag"), authMiddlewareFunc, controller.GetFlag)

	r.PUT("/flags/:flagid", RouteName("update_flag"), authMiddlewareFunc, controller.UpdateFlag)
	r.POST("/flags", RouteName("create_flag"), authMiddlewareFunc, controller.CreateFlag)
	r.DELETE("/flags/:flagid", RouteName("delete_flag"), authMiddlewareFunc, controller.DeleteFlag)
}

func handleNoRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"uptime":    uptime().String(),
		"message":   "OK",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func meHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	u, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": u.(*user.User).UserName,
		"role":     u.(*user.User).Role,
	})
}
func RouteName(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("routeName", name)
		c.Next()
	}
}
