package category

import (
	mockdb "blog-api/db/mock"
	db "blog-api/db/sqlc"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestListAccountsAPI(t *testing.T) {
	n := 5
	// category := randomCategory(t)
	categories := make([]db.Category, n)
	for i := 0; i < n; i++ {
		categories[i] = randomCategory(t)
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",

			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetCategories(gomock.Any()).
					Times(1).
					Return(categories, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCategories(t, recorder.Body, categories)
			},
		},
		// {
		// 	name: "NoAuthorization",

		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			ListAccounts(gomock.Any(), gomock.Any()).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)
		// 	},
		// },
		{
			name: "InternalError",

			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCategories(gomock.Any()).
					Times(1).
					Return([]db.Category{}, sql.ErrConnDone)
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

			url := "/api/v1/category/"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			q := request.URL.Query()
			request.URL.RawQuery = q.Encode()
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
