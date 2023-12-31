package category

import (
	db "blog-api/db/sqlc"
	serverpkg "blog-api/pkg/server"
	"blog-api/util"

	// "fmt"
	"net/http"

	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
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
	NewCategoryHandlers(groupRouter, server.Store, server.Config)
	return nil
}

func NewCategoryHandlers(router *gin.RouterGroup, store db.Store, config util.Config) {
	categoryHandler := NewCategoryUsecase(store, config)
	route := router.Group("/category")
	route.POST("/create", categoryHandler.CreateCategory)
	route.GET("/:id", categoryHandler.FetchCategory)
	route.PUT("/:id", categoryHandler.UpdateCategory)
	route.DELETE("/:id", categoryHandler.DeleteCategory)
	route.GET("/", categoryHandler.ListCategory)

}
