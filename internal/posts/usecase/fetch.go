package posts

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPostRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (usecase *postsusecase) Fetch(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse("error", err.Error()))
		return
	}
	post, err := usecase.Store.GetPostById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse("error", err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("error", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, post)
}
