package tags

import (
	db "blog-api/db/sqlc"
	"blog-api/util"

	"github.com/gin-gonic/gin"
)

type tagsusecase struct {
	store  db.Store
	config util.Config
}

type Tagsusecase interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Fetch(ctx *gin.Context)
	List(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewTagUsecase(store db.Store, config util.Config) Tagsusecase {
	return &tagsusecase{store: store, config: config}
}
