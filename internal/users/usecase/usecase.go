package users

import (
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"

	"github.com/gin-gonic/gin"
)

type userusecase struct {
	store      db.Store
	config     util.Config
	tokenMaker tokenpkg.Maker
}

type Usersusecase interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

func NewUserUsecase(store db.Store, config util.Config, tokenMaker tokenpkg.Maker) Usersusecase {
	return &userusecase{store: store, config: config, tokenMaker: tokenMaker}
}
