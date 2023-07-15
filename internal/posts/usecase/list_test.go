package posts

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"blog-api/util"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestListPostWithCommentsAPI(t *testing.T) {
	n := 5
	posts := make([]db.ListPostWithCommentAndTagsRow, n)
	category := randomCategory(t)
	user := randomUser(t)
	for i := 0; i < n; i++ {

		post := randomPostCommentsTag(t, user, category)
		posts[i] = post
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					ListPostWithCommentAndTags(gomock.Any()).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPosts(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostWithCommentAndTags(gomock.Any()).
					Times(1).
					Return([]db.ListPostWithCommentAndTagsRow{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			url := "/api/v1/blog/"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			q := request.URL.Query()
			request.URL.RawQuery = q.Encode()
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

// ////// list posts with comments and tags based on the selected category
func TestListPostWithCommentsViaCategoryAPI(t *testing.T) {
	n := 5
	posts := make([]db.ListPostbyCategoriesRow, n)
	category := randomCategory(t)
	user := randomUser(t)
	for i := 0; i < n; i++ {

		post := randomPostCommentsTagViaCategory(t, user, category)
		posts[i] = post
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
		categoryId    int32
	}{
		{
			name:       "OK",
			categoryId: category.ID,

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					ListPostbyCategories(gomock.Any(), gomock.Any()).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
			},
			categoryId: category.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPosts(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			categoryId: category.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostbyCategories(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.ListPostbyCategoriesRow{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidParameter",
			categoryId: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostbyCategories(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/api/v1/blog/categories/%d", tc.categoryId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			q := request.URL.Query()
			request.URL.RawQuery = q.Encode()
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

// ////// list posts with comments and tags based on the selected tagd
func TestListPostWithCommentsViaTagAPI(t *testing.T) {
	n := 5
	posts := make([]db.ListPostbyTagRow, n)
	category := randomCategory(t)
	user := randomUser(t)
	for i := 0; i < n; i++ {

		post := randomPostCommentsTagViaTag(t, user, category)
		posts[i] = post
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
		categoryId    int32
	}{
		{
			name:       "OK",
			categoryId: category.ID,

			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					ListPostbyTag(gomock.Any(), gomock.Any()).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
			},
			categoryId: category.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostbyTag(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			categoryId: category.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostbyTag(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.ListPostbyTagRow{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidParameter",
			categoryId: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPostbyTag(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/api/v1/blog/tags/%d", tc.categoryId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			q := request.URL.Query()
			request.URL.RawQuery = q.Encode()
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomPostCommentsTag(t *testing.T, user db.User, category db.Category) (post db.ListPostWithCommentAndTagsRow) {

	post = db.ListPostWithCommentAndTagsRow{
		ID:         int32(util.RandomInt(0, 10000000)),
		Title:      util.RandomString(12),
		Content:    util.RandomString(50),
		AuthorID:   user.ID,
		CategoryID: category.ID,
	}
	return
}

func randomPostCommentsTagViaCategory(t *testing.T, user db.User, category db.Category) (post db.ListPostbyCategoriesRow) {

	post = db.ListPostbyCategoriesRow{
		ID:         int32(util.RandomInt(0, 10000000)),
		Title:      util.RandomString(12),
		Content:    util.RandomString(50),
		AuthorID:   user.ID,
		CategoryID: category.ID,
	}
	return
}

func randomPostCommentsTagViaTag(t *testing.T, user db.User, category db.Category) (post db.ListPostbyTagRow) {

	post = db.ListPostbyTagRow{
		ID:         int32(util.RandomInt(0, 10000000)),
		Title:      util.RandomString(12),
		Content:    util.RandomString(50),
		AuthorID:   user.ID,
		CategoryID: category.ID,
	}
	return
}
