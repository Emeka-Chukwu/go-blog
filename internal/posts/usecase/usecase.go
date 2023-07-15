package posts

import (
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"

	"github.com/gin-gonic/gin"
)

type postsusecase struct {
	Store  db.Store
	Config util.Config
	token  tokenpkg.Maker
}

type Postsusecase interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Fetch(ctx *gin.Context)
	List(ctx *gin.Context)
	ListPostbyTag(ctx *gin.Context)
	ListPostbyCategories(ctx *gin.Context)
	// DeleteCategorye(ctx *gin.Context)
}

func NewPostUsecase(store db.Store, config util.Config, token tokenpkg.Maker) Postsusecase {
	return &postsusecase{Store: store, Config: config}
}
