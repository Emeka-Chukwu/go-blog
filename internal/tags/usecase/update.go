package tags

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (usecase *tagsusecase) Update(ctx *gin.Context) {
	var uriReq getFetchRequest
	var req updateRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	if err := ctx.ShouldBindUri(uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	arg := db.UpdateTagParams{
		Name: req.Name,
	}
	tag, err := usecase.store.UpdateTag(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, util.FormatErrorResponse(pqErr.Code.Name(), err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Error occured", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

type updateRequest struct {
	Name string `json:"name"`
	ID   *int32 `json:"id"`
}
