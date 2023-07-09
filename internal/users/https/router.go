package users

import (
	db "blog-api/db/sqlc"
	usecase "blog-api/internal/users/usecase"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"

	// usecase "eventhuz-api-backend/internal/auths/usecase"

	"github.com/gin-gonic/gin"
)

func NewTagsHandlers(router *gin.RouterGroup, store db.Store, config util.Config, tokenMaker tokenpkg.Maker) {
	handler := usecase.NewUserUsecase(store, config, tokenMaker)
	route := router.Group("/auths")
	route.POST("/signup", handler.Signup)
	route.POST("/login", handler.Login)

}
