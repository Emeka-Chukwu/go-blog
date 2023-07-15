package posts

import (
	db "blog-api/db/sqlc"
	usecase "blog-api/internal/posts/usecase"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"

	"github.com/gin-gonic/gin"
)

func NewPostsHandlers(router *gin.RouterGroup, store db.Store, config util.Config, token tokenpkg.Maker) {
	postHandler := usecase.NewPostUsecase(store, config, token)
	route := router.Group("/blog")
	route.POST("/create", postHandler.Create)
	route.GET("/:id", postHandler.Fetch)
	route.PUT("/:id", postHandler.Update)
	// route.DELETE("/:id", categoryHandler.DeleteCategory)
	route.GET("/", postHandler.List)
	route.GET("/categories/:category_id", postHandler.ListPostbyCategories)
	route.GET("/tags/:tag_id", postHandler.ListPostbyTag)
}
