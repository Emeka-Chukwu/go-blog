package tags

import (
	db "blog-api/db/sqlc"
	usecase "blog-api/internal/tags/usecase"
	"blog-api/util"

	// usecase "eventhuz-api-backend/internal/auths/usecase"

	"github.com/gin-gonic/gin"
)

func NewTagsHandlers(router *gin.RouterGroup, store db.Store, config util.Config) {
	handler := usecase.NewTagUsecase(store, config)
	route := router.Group("/tags")
	route.POST("/create", handler.Create)
	route.GET("/:id", handler.Fetch)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
	route.GET("/", handler.List)
}
