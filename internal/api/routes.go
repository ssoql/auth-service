package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ssoql/auth-service/internal/api/handlers"
	"github.com/ssoql/auth-service/internal/infrastructure"
	"github.com/ssoql/auth-service/internal/infrastructure/db"
	"github.com/ssoql/auth-service/internal/use_cases"
)

func RegisterRoutes(router *gin.Engine, dbClient *db.ClientDB) {
	dbRead := infrastructure.NewUserReadRepository(dbClient)
	dbWrite := infrastructure.NewUserWriteRepository(dbClient)

	saveUserUseCase := use_cases.NewSaveUserUseCase(dbRead, dbWrite)
	loginUseCase := use_cases.NewLoginUseCase(dbRead)
	validateUseCase := use_cases.NewValidateUseCase()

	userCreateHandler := handlers.NewUserCreateHandler(saveUserUseCase)
	loginHandler := handlers.NewLoginHandler(loginUseCase)
	validateHandler := handlers.NewValidateHandler(validateUseCase)

	router.POST("/users", userCreateHandler.Handle) // will be extracted to separate service
	router.POST("/token", loginHandler.Handle)
	router.POST("/token/validate", validateHandler.Handle)
}
