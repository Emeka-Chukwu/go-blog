package tags

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *tagsusecase) Delete(ctx *gin.Context) {
	var req getFetchRequest
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	err := usecase.store.DeleteTag(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err.Error()))
			return
		}
		ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
