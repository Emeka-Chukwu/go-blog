package category

import (
	"blog-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *categoryusecase) ListCategory(ctx *gin.Context) {
	data, err := usecase.Store.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, data)
}
