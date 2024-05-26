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
)

var startTime time.Time // Used to calculate uptime
var users map[string]User

func init() {
	startTime = time.Now()
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	users = make(map[string]User)
	users["admin"] = User{
		UserName: "admin",
		Role:     "admin",
	}
	users["test"] = User{
		UserName: "test",
		Role:     "user",
	}

}
func uptime() time.Duration {
	return time.Since(startTime)
}

func main() {
	engine := gin.Default()
	// the jwt middleware
	adminMiddleware, err := jwt.New(newAuthMiddleware("admin"))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	userMiddleware, err := jwt.New(newAuthMiddleware("user"))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// register middleware
	engine.Use(handlerMiddleWare(adminMiddleware))

	// register route
	registerRoute(engine, adminMiddleware, userMiddleware)

	// start http server
	if err = http.ListenAndServe(":"+port, engine); err != nil {
		log.Fatal(err)
	}
}

func registerRoute(r *gin.Engine, adminHandler *jwt.GinJWTMiddleware, userHandler *jwt.GinJWTMiddleware) {
	r.GET("/health-check", healthHandler)
	r.POST("/login", adminHandler.LoginHandler)
	r.NoRoute(adminHandler.MiddlewareFunc(), handleNoRoute())

	auth := r.Group("/", adminHandler.MiddlewareFunc())
	auth.GET("/refresh_token", adminHandler.RefreshHandler)

	r.GET("/me", userHandler.MiddlewareFunc(), meHandler)

	params, paramErr := envs.GetConnectionParams()
	if len(paramErr) > 0 {
		errMsg := "Param errors have been found:\n"
		for _, err := range paramErr {
			errMsg += fmt.Sprintf("  - %v\n", err.Error())
		}
		log.Fatal(errMsg)
	}

	// repository := storage.NewMemoryRepository()
	db, err := flags.NewFlagMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

	repository := storage.NewRedisClient(params.RedisHost, params.RedisPort)
	service := flags.NewFlagService(repository, db)
	controller := NewFlagController(service)

	controller.RegisterRoutes(auth)

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
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"role":     user.(*User).Role,
	})
}
