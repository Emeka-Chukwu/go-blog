package posts

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *postsusecase) List(ctx *gin.Context) {
	posts, err := usecase.Store.ListPostWithCommentAndTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (usecase *postsusecase) ListPostbyCategories(ctx *gin.Context) {
	var req getCategoryId
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse("error", err.Error()))
		return
	}
	arg, err := usecase.Store.ListPostbyCategories(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, arg)
}

func (usecase *postsusecase) ListPostbyTag(ctx *gin.Context) {
	var req getTagId
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse("error", err.Error()))
		return
	}
	arg, err := usecase.Store.ListPostbyTag(ctx, sql.NullInt32{Valid: true, Int32: req.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, arg)
}

type getCategoryId struct {
	ID int32 `uri:"category_id" binding:"required"`
}

type getTagId struct {
	ID int32 `uri:"tag_id" binding:"required"`
}
