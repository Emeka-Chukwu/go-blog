package tags

import (
	"blog-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (usecase *tagsusecase) List(ctx *gin.Context) {
	tags, err := usecase.store.GetTags(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	ctx.JSON(http.StatusOK, tags)
}
