package users

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (usecase *userusecase) Signup(ctx *gin.Context) {
	var req signupCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	ID := int32(util.RandomInt(0, 2000000))
	if req.ID != nil {
		ID = *req.ID
	}
	arg := db.CreateUserParams{
		ID:       ID,
		Username: req.Username,
		Email:    req.Email,
		Role:     "user",
		Password: hashedPassword,
	}
	user, err := usecase.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, util.FormatErrorResponse("email already exist", err.Error()))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	ctx.JSON(http.StatusCreated, newUserResponse(user))
}

type signupCreateRequest struct {
	ID       *int32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
