package posts

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestUpdatePostAPI(t *testing.T) {
	category := randomCategory(t)
	tag := randomTag(t)
	user := randomUser(t)
	post, _ := randomPostTag(t, user, category, tag)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"id":          post.ID,
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
				"author_id":   post.AuthorID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdatePostParams{
					ID:         post.ID,
					Title:      sql.NullString{Valid: true, String: post.Title},
					Content:    sql.NullString{Valid: true, String: post.Content},
					CategoryID: sql.NullInt32{Valid: true, Int32: post.CategoryID},
					AuthorID:   post.AuthorID,
				}

				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(post, nil)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPost(t, recorder.Body, post)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"id":          post.ID,
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
				"author_id":   post.AuthorID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"id":          post.ID,
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
				"author_id":   post.AuthorID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Post{}, sql.ErrConnDone)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidData",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"category_id": "kkkk",
				"tag_id":      tag.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "RecordNotFound",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Post{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ExistingBlogTitle",
			body: gin.H{
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Post{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		// Return(db.User{}, &pq.Error{Code: "23505"})
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/api/v1/blog/%d", post.ID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
