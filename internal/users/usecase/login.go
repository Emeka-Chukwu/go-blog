package users

import (
	db "blog-api/db/sqlc"
	"blog-api/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (usecase *userusecase) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.FormatErrorResponse(err.Error(), err.Error()))
		return
	}
	foundUser, err := usecase.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.FormatErrorResponse(err.Error(), err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}

	err = util.CheckPassword(req.Password, foundUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	accessToken, payload, err := usecase.tokenMaker.CreateToken(
		foundUser.ID, foundUser.Role,
		usecase.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.FormatErrorResponse(err.Error(), err))
		return
	}
	resp := loginUserResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: payload.ExpiredAt,
		User:                 newUserResponse(foundUser),
	}
	ctx.JSON(http.StatusOK, resp)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	AccessToken          string       `json:"access_token"`
	AccessTokenExpiresAt time.Time    `json:"access_token_expires_at"`
	User                 UserResponse `json:"user"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdatedAt,
	}
}

type UserResponse struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
