package serverpkg

import (
	db "blog-api/db/sqlc"
	"blog-api/util"

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
	// server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
