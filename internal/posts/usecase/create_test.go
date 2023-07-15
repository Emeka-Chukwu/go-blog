package posts

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreatePostAPI(t *testing.T) {
	category := randomCategory(t)
	tag := randomTag(t)
	user := randomUser(t)
	post, postTag := randomPostTag(t, user, category, tag)

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
				arg := db.CreatePostParams{
					ID:         post.ID,
					Title:      post.Title,
					Content:    post.Content,
					CategoryID: post.CategoryID,
					AuthorID:   post.AuthorID,
				}
				argPostTag := db.CreateTagsToPostParams{
					PostID: sql.NullInt32{Valid: true, Int32: post.ID},
					TagID:  sql.NullInt32{Valid: true, Int32: tag.ID},
				}
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(post, nil)
				store.EXPECT().
					CreateTagsToPost(gomock.Any(), gomock.Eq(argPostTag)).
					Times(1).
					Return(postTag, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
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
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateTagsToPost(gomock.Any(), gomock.Any()).
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
					CreatePost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Post{}, sql.ErrConnDone)

				store.EXPECT().
					CreateTagsToPost(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidData",
			body: gin.H{
				"id":          "0000",
				"title":       post.Title,
				"content":     post.Content,
				"category_id": post.CategoryID,
				"tag_id":      tag.ID,
				"author_id":   "oo",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
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

			url := "/api/v1/blog/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomCategory(t *testing.T) (category db.Category) {
	category = db.Category{
		Name: sql.NullString{Valid: true, String: util.RandomString(12)},
		ID:   int32(util.RandomInt(1, 1000000000)),
	}
	return
}

func randomTag(t *testing.T) (tag db.Tag) {
	tag = db.Tag{
		Name: util.RandomString(12),
		ID:   int32(util.RandomInt(1, 1000000000)),
	}
	return
}

func randomUser(t *testing.T) (user db.User) {
	password := util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	user = db.User{
		ID:       int32(util.RandomInt(0, 10000000)),
		Username: util.RandomUsername(),
		Password: hashedPassword,
		Email:    util.RandomEmail(),
		Role:     "user",
	}
	return
}
func randomPostTag(t *testing.T, user db.User, category db.Category, tag db.Tag) (post db.Post, tagPost db.PostTag) {

	post = db.Post{
		ID:         int32(util.RandomInt(0, 10000000)),
		Title:      util.RandomString(12),
		Content:    util.RandomString(50),
		AuthorID:   user.ID,
		CategoryID: category.ID,
	}
	tagPost = db.PostTag{
		ID:     int32(util.RandomInt(0, 10000000)),
		PostID: sql.NullInt32{Valid: true, Int32: post.ID},
		TagID:  sql.NullInt32{Valid: true, Int32: tag.ID},
	}
	return
}

func requireBodyMatchPost(t *testing.T, body *bytes.Buffer, post db.Post) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotPost db.Post
	err = json.Unmarshal(data, &gotPost)
	require.NoError(t, err)
	require.Equal(t, post.Title, gotPost.Title)
	require.Equal(t, post.Content, gotPost.Content)
	require.Equal(t, post.AuthorID, gotPost.AuthorID)
	require.Equal(t, post.CategoryID, gotPost.CategoryID)
	require.NotZero(t, post.CategoryID, gotPost.CategoryID)

}

func requireBodyMatchCategories(t *testing.T, body *bytes.Buffer, categories []db.Category) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotCategories []db.Category
	err = json.Unmarshal(data, &gotCategories)
	require.NoError(t, err)
	require.Equal(t, categories, gotCategories)

}
