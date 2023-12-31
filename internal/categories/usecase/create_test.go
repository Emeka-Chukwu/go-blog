package category

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	"blog-api/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateCategoryAPI(t *testing.T) {
	categoryModel := randomCategory(t)
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
				"name": categoryModel.Name.String,
				"id":   categoryModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateCategoryParams{
					Name: categoryModel.Name,
					ID:   categoryModel.ID,
				}
				fmt.Println(arg)
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(categoryModel, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchCategory(t, recorder.Body, categoryModel)
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
				"name": categoryModel.Name.String,
				"id":   categoryModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Category{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidData",
			body: gin.H{
				"name": categoryModel.Name,
				"id":   categoryModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "DuplicateRecord",
			body: gin.H{
				"name": categoryModel.Name.String,
				"id":   categoryModel.ID,
			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCategory(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Category{}, &pq.Error{Code: "23505"})
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

			url := "/api/v1/category/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			// tc.setupAuth(t, request, server.tokenMaker)
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

func requireBodyMatchCategory(t *testing.T, body *bytes.Buffer, category db.Category) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotCategory db.Category
	err = json.Unmarshal(data, &gotCategory)
	require.NoError(t, err)
	require.Equal(t, category.Name, gotCategory.Name)
	require.NotZero(t, gotCategory.ID)

}

func requireBodyMatchCategories(t *testing.T, body *bytes.Buffer, categories []db.Category) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	var gotCategories []db.Category
	err = json.Unmarshal(data, &gotCategories)
	require.NoError(t, err)
	require.Equal(t, categories, gotCategories)

}
