package tags

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getFetchRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (usecase *tagsusecase) Fetch(ctx *gin.Context) {
	var req getFetchRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err))
		return
	}
	tag, err := usecase.store.GetTagId(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	ctx.JSON(http.StatusOK, tag)
}
