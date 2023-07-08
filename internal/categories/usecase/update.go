package category

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (usecase *categoryusecase) UpdateCategory(ctx *gin.Context) {

	var req updateCategoryRequest
	var getReq getCategoryRequest
	if err := ctx.ShouldBindUri(&getReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	arg := db.UpdateCategoryParams{
		Name: sql.NullString{Valid: true, String: req.Name},
		ID:   getReq.ID,
	}
	category, err := usecase.Store.UpdateCategory(ctx, arg)
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
	// rsp := u.DataResponse{
	// 	Message: "Event post updated successfully",
	// 	Data:    eventModel,
	// }
	ctx.JSON(http.StatusCreated, category)

}

type updateCategoryRequest struct {
	Name string `json:"name"`
	ID   *int32 `json:"id"`
}
