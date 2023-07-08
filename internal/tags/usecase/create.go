package tags

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (usecase *tagsusecase) Create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err))
		return
	}
	Id := int32(util.RandomInt(1, 1000000000))
	if req.ID != nil {
		Id = *req.ID
	}
	arg := db.CreateTagsParams{
		ID:   Id,
		Name: req.Name,
	}
	categoryResp, err := usecase.store.CreateTags(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, util.FormatErrorResponse("tags already exist", err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	ctx.JSON(http.StatusCreated, categoryResp)

}

type createRequest struct {
	Name string `json:"name"`
	ID   *int32 `json:"id"`
}
