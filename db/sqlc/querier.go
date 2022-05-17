// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreatePokemon(ctx context.Context, arg CreatePokemonParams) (Pokemon, error)
	DeletePokemon(ctx context.Context, id int64) error
	GetPokemon(ctx context.Context, id int64) (Pokemon, error)
	ListPokemons(ctx context.Context, arg ListPokemonsParams) ([]Pokemon, error)
}

var _ Querier = (*Queries)(nil)