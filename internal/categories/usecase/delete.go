package category

import (
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *categoryusecase) DeleteCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	err := usecase.Store.DeleteCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	// rsp := resp.DataResponse{
	// 	Message: "Event deleted successfully",
	// 	Data:    nil,
	// }
	ctx.JSON(http.StatusOK, nil)
}
