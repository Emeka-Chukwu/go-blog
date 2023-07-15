package posts

import (
	db "blog-api/db/sqlc"
	helper "blog-api/internal/posts/helpers"
	"blog-api/util"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (usecase *postsusecase) Update(ctx *gin.Context) {
	var req requestUpdatePostParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var getReq getPostRequest
	if err := ctx.ShouldBindUri(&getReq); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	payload, exist := helper.GetAuthenticatedPayload(ctx)
	if !exist {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Internal error", ""))
		return
	}
	arg := db.UpdatePostParams{
		Title:      sql.NullString{Valid: req.Title != nil, String: *req.Title},
		Content:    sql.NullString{Valid: req.Content != nil, String: *req.Content},
		CategoryID: sql.NullInt32{Valid: req.CategoryID != nil, Int32: *req.CategoryID},
		AuthorID:   payload.UserID,
		ID:         getReq.ID,
	}
	post, err := usecase.Store.UpdatePost(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, util.FormatErrorResponse(pqErr.Code.Name(), err.Error()))
				return
			}
		} else if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Error occured", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, post)

}

type requestUpdatePostParams struct {
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	CategoryID *int32  `json:"category_id"`
	TagID      *int32  `json:"tag_id"`
}
