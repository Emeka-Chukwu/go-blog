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

	Id := int32(util.RandomInt(1, 1000000000))
	if req.ID != nil {
		Id = *req.ID
	}
	arg := db.CreateCategoryParams{
		ID:   Id,
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
	// rsp := util.DataResponse{
	// 	Message: "category record saved successfully",
	// 	Data:    categoryResp,
	// }
	ctx.JSON(http.StatusCreated, categoryResp)

}

type createCategoryRequest struct {
	Name string `json:"name"`
	ID   *int32 `json:"id"`
}

func categoryResponse(cat db.Category) {

}
