package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(15 * time.Second)
		panic("test")
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	middleware.AddHttpMiddleware(router)
	api.RegisterRoutes(router, dbClient)

	// Setup HTTP server
	srv := &http.Server{
		Addr:    env.GetPort(),
		Handler: router,
	}

	// Run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server started on http://localhost" + env.GetPort())
	gracefulShutdown(srv, 5*time.Second)
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

// gracefulShutdown listens for interrupt signals and gracefully shuts down the server
func gracefulShutdown(srv *http.Server, timeout time.Duration) {
	// Create a channel to listen for termination signals
	quit := make(chan os.Signal, 1)

	// SIGINT (Ctrl+C), SIGTERM (docker stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive the signal
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting gracefully")
}
