package category

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (category *categoryusecase) CreateCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err))
		return
	}
	arg := db.CreateCategoryParams{
		ID:   int32(util.RandomInt(1, 1000000000)),
		Name: sql.NullString{Valid: true, String: req.Name},
	}
	categoryResp, err := category.Store.CreateCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, util.FormatErrorResponse("category already exist", err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	rsp := util.DataResponse{
		Message: "User record saved successfully",
		Data:    categoryResp,
	}
	ctx.JSON(http.StatusCreated, rsp)

}

type createCategoryRequest struct {
	Name string `json:"name"`
}
