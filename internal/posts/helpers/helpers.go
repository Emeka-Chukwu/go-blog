package posts

import (
	tokenpkg "blog-api/pkg/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func GetAuthenticatedPayload(ctx *gin.Context) (payload *tokenpkg.Payload, exist bool) {
	author, exist := ctx.Get(authorizationPayloadKey)
	if !exist {
		return
	}
	payload, exist = author.(*tokenpkg.Payload)
	return
}
