package app

import (
	db "blog-api/db/sqlc"
	categoryhandler "blog-api/internal/categories/https"
	posthandler "blog-api/internal/posts/https"
	taghandler "blog-api/internal/tags/https"
	userhandler "blog-api/internal/users/https"
	"blog-api/middleware"
	serverpkg "blog-api/pkg/server"
	"blog-api/util"
	"net/http"

	// db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

func InitializeServer(config util.Config, store db.Store) (*serverpkg.Server, error) {
	data, err := serverpkg.NewServer(config, store)
	SetupRouter(data)
	return data, err
}

func SetupRouter(server *serverpkg.Server) {
	router := gin.Default()
	server.Router = router

	groupRouter := router.Group("/api/v1")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "app is unning fine at" + server.Config.HTTPServerAddress})
	})
	userhandler.NewTagsHandlers(groupRouter, server.Store, server.Config, server.TokenMaker)
	groupRouter.Use(middleware.AuthMiddleware(server.TokenMaker, server.Config))
	categoryhandler.NewCategoryHandlers(groupRouter, server.Store, server.Config)
	taghandler.NewTagsHandlers(groupRouter, server.Store, server.Config)
	posthandler.NewPostsHandlers(groupRouter, server.Store, server.Config, server.TokenMaker)
	// userHandler.NewUserHandlers(groupRouter, server.store, server.taskDistributor, server.tokenMaker, server.config)

}
