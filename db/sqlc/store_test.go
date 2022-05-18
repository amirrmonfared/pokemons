package db

import (
	"context"
	"testing"

	"github.com/amirrmonfared/pokemons/util"
	"github.com/stretchr/testify/require"
)

func TestImportPokemon(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan CreatePokemonResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.ImportPokemon(context.Background(), CreatePokemonParams{
				Name:       util.RandomString(4),
				Type1:      util.RandomString(4),
				Type2:      util.RandomString(4),
				Total:      util.RandomInt(200, 500),
				Hp:         util.RandomInt(1, 200),
				Attack:     util.RandomInt(1, 100),
				Defense:    util.RandomInt(1, 100),
				SpAtk:      util.RandomInt(1, 100),
				SpDef:      util.RandomInt(1, 100),
				Speed:      util.RandomInt(1, 100),
				Generation: util.RandomInt(1, 100),
				Legendary:  util.RandomBool(),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
	}
}
