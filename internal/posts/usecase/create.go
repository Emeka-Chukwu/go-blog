package posts

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"database/sql"
	"net/http"

	helper "blog-api/internal/posts/helpers"

	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func (usecase *postsusecase) Create(ctx *gin.Context) {
	var req requestPostParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse("invalid data", err.Error()))
		return
	}
	payload, exist := helper.GetAuthenticatedPayload(ctx)
	if !exist {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Internal error", ""))
		return
	}
	postID := util.RandomInt(0, 488498390)
	if req.ID != nil {
		postID = *req.ID
	}
	postArg := db.CreatePostParams{
		ID:         int32(postID),
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   payload.UserID,
		CategoryID: req.CategoryID,
	}
	post, err := usecase.Store.CreatePost(ctx, postArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Internal error", err.Error()))
		return
	}
	postTagArg := db.CreateTagsToPostParams{
		PostID: sql.NullInt32{Valid: true, Int32: post.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: req.TagID},
	}
	_, err = usecase.Store.CreateTagsToPost(ctx, postTagArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse("Internal error", err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, post)

}

type requestPostParams struct {
	ID         *int64 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int32  `json:"category_id"`
	TagID      int32  `json:"tag_id"`
}
