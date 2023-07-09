package users

import (
	db "blog-api/db/sqlc"
	serverpkg "blog-api/pkg/server"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"
	"fmt"

	// "fmt"
	"net/http"

	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func newTestServer(t *testing.T, store db.Store) *serverpkg.Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := serverpkg.NewServer(config, store)
	err = SetupRouter(server)
	if err != nil {
		return &serverpkg.Server{}
	}
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func SetupRouter(server *serverpkg.Server) error {
	router := gin.Default()
	server.Router = router

	groupRouter := router.Group("/api/v1")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "app is unning fine at" + server.Config.HTTPServerAddress})
	})
	NewTagHandlers(groupRouter, server.Store, server.Config, server.TokenMaker)
	return nil
}

func NewTagHandlers(router *gin.RouterGroup, store db.Store, config util.Config, tokenMaker tokenpkg.Maker) {
	handler := NewUserUsecase(store, config, tokenMaker)
	route := router.Group("/auths")
	route.POST("/signup", handler.Signup)
	route.POST("/login", handler.Login)
}

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker tokenpkg.Maker,
	authorizationType string,
	username int32,
	duration time.Duration,
	role string,
) {
	token, payload, err := tokenMaker.CreateToken(username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}
