package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ssoql/auth-service/config"
	"github.com/ssoql/auth-service/internal/api"
	"github.com/ssoql/auth-service/internal/api/middleware"
	"github.com/ssoql/auth-service/internal/infrastructure/db"

	"github.com/gin-contrib/cors"
)

func main() {
	env := config.NewEnv()

	dbClient, err := initializeDB(env)
	if err != nil {
		panic(err)
	}

	router := createRouter()

	middleware.AddHttpMiddleware(router)
	api.RegisterRoutes(router, dbClient)

	if err := router.Run(env.GetPort()); err != nil {
		panic(err)
	}
}

func createRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	return router
}

func initializeDB(env *config.Env) (*db.ClientDB, error) {
	dbConfig := db.DatabaseOptions{
		DatabaseUser:     env.GetDbUser(),
		DatabasePassword: env.GetDbPassword(),
		DatabaseSchema:   env.GetDbSchema(),
		DatabasePort:     env.GetDbPort(),
		DatabaseHost:     env.GetDbHost(),
	}

	fmt.Printf("DB conf:\n %v", dbConfig)

	return db.InitializeDB(dbConfig)
}
