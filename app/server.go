package app

import (
	db "blog-api/db/sqlc"
	categoryhandler "blog-api/internal/categories/https"
	"blog-api/util"
	"net/http"

	// db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config util.Config
	Store  db.Store
	Router *gin.Engine
}

//// server serves out http request for our backend service

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{Store: store, Config: config}
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	server.Router = router

	groupRouter := router.Group("/api/v1")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "app is unning fine at" + server.Config.HTTPServerAddress})
	})
	categoryhandler.NewCategoryHandlers(groupRouter, server.Store, server.Config)
	// userHandler.NewUserHandlers(groupRouter, server.store, server.taskDistributor, server.tokenMaker, server.config)

}
