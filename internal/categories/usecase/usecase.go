package category

import (
	db "blog-api/db/sqlc"
	"blog-api/util"

	"github.com/gin-gonic/gin"
)

type categoryusecase struct {
	Store  db.Store
	Config util.Config
}

type Categoryusecase interface {
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	FetchCategory(ctx *gin.Context)
	ListCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
}

func NewCategoryUsecase(store db.Store, config util.Config) Categoryusecase {
	return &categoryusecase{Store: store, Config: config}
}
