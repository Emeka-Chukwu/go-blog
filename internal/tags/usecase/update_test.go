package tags

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestUpdateTagAPI(t *testing.T) {
	tag := randomTag(t)
	testCases := []struct {
		name  string
		body  gin.H
		tagId int32
		// setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": tag.Name,
				"id":   tag.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			tagId: tag.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateTagParams{
					Name: tag.Name,
					ID:   tag.ID,
				}
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(tag, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTag(t, recorder.Body, tag)
			},
		},
		// {
		// 	name: "NoAuthorization",
		// 	body: gin.H{
		// 		"currency": account.Currency,
		// 	},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			CreateAccount(gomock.Any(), gomock.Any()).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		// 	},
		// },
		{
			name:  "DuplicateRecord",
			tagId: tag.ID,
			body: gin.H{
				"name": tag.Name,
				"id":   tag.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Tag{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			tagId: tag.ID,
			body: gin.H{
				"name": tag.Name,
				"id":   tag.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Tag{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "InvalidData",
			tagId: tag.ID,
			body: gin.H{
				"name": 88,
				"id":   "",
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:  "RecordNotFound",
			tagId: tag.ID,
			body: gin.H{
				"name": tag.Name,
				"id":   888788,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Tag{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "Invalidparams",
			tagId: 0,
			body: gin.H{
				"name": tag.Name,
				"id":   tag.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateTag(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/api/v1/tags/%d", tc.tagId)
			fmt.Println(url, "kkkkkk")
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)
			// tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
