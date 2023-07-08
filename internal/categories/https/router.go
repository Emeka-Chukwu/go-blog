package category

import (
	db "blog-api/db/sqlc"
	usecase "blog-api/internal/categories/usecase"
	"blog-api/util"

	// usecase "eventhuz-api-backend/internal/auths/usecase"

	"github.com/gin-gonic/gin"
)

func NewCategoryHandlers(router *gin.RouterGroup, store db.Store, config util.Config) {
	categoryHandler := usecase.NewCategoryUsecase(store, config)
	route := router.Group("/category")
	route.POST("/create", categoryHandler.CreateCategory)
	route.GET("/:id", categoryHandler.FetchCategory)
	route.PUT("/:id", categoryHandler.UpdateCategory)
	route.DELETE("/:id", categoryHandler.DeleteCategory)
	route.GET("/", categoryHandler.ListCategory)
}
