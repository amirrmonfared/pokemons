package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/amirrmonfared/pokemons/db/mock"
	db "github.com/amirrmonfared/pokemons/db/sqlc"
	"github.com/amirrmonfared/pokemons/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomPokemon() db.Pokemon {
	return db.Pokemon{
		Name:       util.RandomString(4),
		Type1:      util.RandomString(4),
		Type2:      util.RandomString(4),
		Total:      util.RandomInt(200, 400),
		Hp:         util.RandomInt(1, 200),
		Attack:     util.RandomInt(1, 100),
		Defense:    util.RandomInt(1, 100),
		SpAtk:      util.RandomInt(1, 100),
		SpDef:      util.RandomInt(1, 100),
		Speed:      util.RandomInt(1, 100),
		Generation: util.RandomInt(1, 100),
		Legendary:  util.RandomBool(),
	}
}

func TestListPokemonsAPI(t *testing.T) {

	n := 10
	pokemons := make([]db.Pokemon, n)
	for i := 0; i < n; i++ {
		pokemons[i] = randomPokemon()
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListPokemonsParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListPokemons(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(pokemons, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPokemons(t, recorder.Body, pokemons)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListPokemons(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Pokemon{}, sql.ErrConnDone)
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

			url := "/pokemon"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchPokemons(t *testing.T, body *bytes.Buffer, pokemons []db.Pokemon) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotPokemons []db.Pokemon
	err = json.Unmarshal(data, &gotPokemons)
	require.NoError(t, err)
	require.Equal(t, pokemons, gotPokemons)
}
