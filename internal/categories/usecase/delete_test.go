package category

import (
	mockdb "blog-api/db/mock"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteCategoryAPI(t *testing.T) {

	category := randomCategory(t)

	testCases := []struct {
		name       string
		categoryID int32
		// setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "Ok",
			categoryID: category.ID,
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteCategory(gomock.Any(), gomock.Eq(category.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},

		// {
		// 	name:      "NoAuthorization",
		// 	accountID: account.ID,
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			GetAccount(gomock.Any(), gomock.Any()).
		// 			Times(0).
		// 			Return(account, nil)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusUnauthorized, recorder.Code)

		// 	},
		// },
		{
			name:       "NotFound",
			categoryID: category.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteCategory(gomock.Any(), gomock.Eq(category.ID)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
		},
		{
			name:       "InternalError",
			categoryID: category.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteCategory(gomock.Any(), gomock.Eq(category.ID)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
		},

		{
			name:       "InvalidID",
			categoryID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCategoryById(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
			// setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			// 	addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			// },
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			/// build stubs
			tc.buildStubs(store)
			////  Start the test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/category/%d", tc.categoryID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			// tc.setupAuth(t, request, server.tokenMaker)
			server.Router.ServeHTTP(recorder, request)
			////check for reponse
			tc.checkResponse(t, recorder)
		})
	}
}
