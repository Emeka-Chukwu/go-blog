package serverpkg

import (
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config     util.Config
	Store      db.Store
	Router     *gin.Engine
	TokenMaker tokenpkg.Maker
}

//// server serves out http request for our backend service

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := tokenpkg.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{Store: store, Config: config, TokenMaker: tokenMaker}
	// server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
