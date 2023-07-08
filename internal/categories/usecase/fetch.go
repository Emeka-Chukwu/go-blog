package category

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (category *categoryusecase) FetchCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse("Invalid data", err.Error()))
		return
	}
	data, err := category.Store.GetCategoryById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, data)
}
