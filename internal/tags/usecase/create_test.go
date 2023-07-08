package tags

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	"blog-api/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateTagAPI(t *testing.T) {
	tagModel := randomTag(t)
	testCases := []struct {
		name string
		body gin.H
		// setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": tagModel.Name,
				"id":   tagModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateTagsParams{
					Name: tagModel.Name,
					ID:   tagModel.ID,
				}
				store.EXPECT().
					CreateTags(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(tagModel, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchTag(t, recorder.Body, tagModel)
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
			name: "InternalError",
			body: gin.H{
				"name": tagModel.Name,
				"id":   tagModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTags(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Tag{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidData",
			body: gin.H{
				"name": 44,
				"id":   tagModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTags(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "DuplicateRecord",
			body: gin.H{
				"name": tagModel.Name,
				"id":   tagModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTags(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Tag{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
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

			url := "/api/v1/tags/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			// tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomTag(t *testing.T) (tag db.Tag) {
	tag = db.Tag{
		Name: util.RandomString(12),
		ID:   int32(util.RandomInt(1, 1000000000)),
	}
	return
}

func requireBodyMatchTag(t *testing.T, body *bytes.Buffer, tag db.Tag) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotTag db.Tag
	err = json.Unmarshal(data, &gotTag)
	require.NoError(t, err)
	require.Equal(t, tag.Name, gotTag.Name)
	require.NotZero(t, gotTag.ID)

}

func requireBodyMatchTags(t *testing.T, body *bytes.Buffer, tags []db.Tag) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotTags []db.Tag
	err = json.Unmarshal(data, &gotTags)
	require.NoError(t, err)
	require.Equal(t, tags, gotTags)

}
