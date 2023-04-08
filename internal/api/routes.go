package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoql/auth-service/internal/api/handlers"
	"github.com/ssoql/auth-service/internal/infrastructure"
	"github.com/ssoql/auth-service/internal/infrastructure/db"
	"github.com/ssoql/auth-service/internal/usecases"
)

func RegisterRoutes(router *gin.Engine, dbClient *db.ClientDB) {
	dbRead := infrastructure.NewUserReadRepository(dbClient)
	dbWrite := infrastructure.NewUserWriteRepository(dbClient)

	saveUserUseCase := usecases.NewSaveUserUseCase(dbRead, dbWrite)
	loginUseCase := usecases.NewLoginUseCase(dbRead)
	validateUseCase := usecases.NewValidateUseCase()

	userCreateHandler := handlers.NewUserCreateHandler(saveUserUseCase)
	loginHandler := handlers.NewLoginHandler(loginUseCase)
	validateHandler := handlers.NewValidateHandler(validateUseCase)

	router.POST("/user", userCreateHandler.Handle)
	router.POST("/login", loginHandler.Handle)
	router.POST("/validate", validateHandler.Handle)
}
