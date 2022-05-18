package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/pokemons/util"
	"github.com/stretchr/testify/require"
)

func createRandomPokemon(t *testing.T) Pokemon {

	arg := CreatePokemonParams{
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

	pokemon, err := testQueries.CreatePokemon(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pokemon)

	require.Equal(t, arg.Name, pokemon.Name)
	require.Equal(t, arg.Type1, pokemon.Type1)
	require.Equal(t, arg.Type2, pokemon.Type2)
	require.Equal(t, arg.Total, pokemon.Total)
	require.Equal(t, arg.Hp, pokemon.Hp)
	require.Equal(t, arg.Attack, pokemon.Attack)
	require.Equal(t, arg.Defense, pokemon.Defense)
	require.Equal(t, arg.SpAtk, pokemon.SpAtk)
	require.Equal(t, arg.SpDef, pokemon.SpDef)
	require.Equal(t, arg.Speed, pokemon.Speed)
	require.Equal(t, arg.Generation, pokemon.Generation)
	require.Equal(t, arg.Legendary, pokemon.Legendary)

	return pokemon
}

func TestCreatePokemon(t *testing.T) {
	createRandomPokemon(t)
}

func TestGetPokemon(t *testing.T) {
	pokemon1 := createRandomPokemon(t)
	pokemon2, err := testQueries.GetPokemon(context.Background(), pokemon1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pokemon2)

	require.Equal(t, pokemon1.ID, pokemon2.ID)
	require.Equal(t, pokemon1.Name, pokemon2.Name)
	require.Equal(t, pokemon1.Type1, pokemon2.Type1)
	require.Equal(t, pokemon1.Type2, pokemon2.Type2)
	require.Equal(t, pokemon1.Hp, pokemon2.Hp)
	require.Equal(t, pokemon1.Attack, pokemon2.Attack)
	require.Equal(t, pokemon1.Defense, pokemon2.Defense)
	require.Equal(t, pokemon1.SpAtk, pokemon2.SpAtk)
	require.Equal(t, pokemon1.SpDef, pokemon2.SpDef)
	require.Equal(t, pokemon1.Speed, pokemon2.Speed)
	require.Equal(t, pokemon1.Generation, pokemon2.Generation)
	require.WithinDuration(t, pokemon1.CreatedAt, pokemon2.CreatedAt, time.Second)
}

func TestDeletePokemon(t *testing.T) {
	pokemon1 := createRandomPokemon(t)
	err := testQueries.DeletePokemon(context.Background(), pokemon1.ID)
	require.NoError(t, err)

	pokemon2, err := testQueries.GetPokemon(context.Background(), pokemon1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pokemon2)
}

func TestListPokemons(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPokemon(t)
	}

	arg := ListPokemonsParams{
		Limit:  5,
		Offset: 0,
	}

	pokemons, err := testQueries.ListPokemons(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pokemons)

	for _, pokemon := range pokemons {
		require.NotEmpty(t, pokemon)
	}
}
