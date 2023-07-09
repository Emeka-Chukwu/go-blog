package middleware

import (
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"
	"errors"

	"fmt"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker tokenpkg.Maker, config util.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifiyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		tokenDiff := payload.ExpiredAt.Sub(payload.IssuedAt).Minutes() + 0.005
		if (tokenDiff - config.AccessTokenDuration.Minutes()) > 1 {
			err := errors.New("invalid token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
