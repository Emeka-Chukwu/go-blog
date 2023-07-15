package posts

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	tokenpkg "blog-api/pkg/token"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFetchPostAPI(t *testing.T) {
	category := randomCategory(t)
	tag := randomTag(t)
	user := randomUser(t)
	post, _ := randomPostTag(t, user, category, tag)

	testCases := []struct {
		postID        int32
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			postID: post.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostById(gomock.Any(), gomock.Eq(post.ID)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "NoAuthorization",
			postID: post.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostById(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			postID: post.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Post{}, sql.ErrConnDone)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InvalidData",
			postID: -0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostById(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "RecordNotFound",
			postID: post.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker tokenpkg.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute, "user")
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetPostById(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Post{}, sql.ErrNoRows)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
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

			url := fmt.Sprintf("/api/v1/blog/%d", tc.postID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
