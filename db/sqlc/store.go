package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	ImportPokemon(ctx context.Context, arg CreatePokemonParams) (CreatePokemonResult, error)
}

// SQLStore provides all functions to execute db queries and transaction
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreatePokemonResult struct {
	Pokemon Pokemon `json:"pokemon"`
}

func (store *SQLStore) ImportPokemon(ctx context.Context, arg CreatePokemonParams) (CreatePokemonResult, error) {
	var result CreatePokemonResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Pokemon, err = q.CreatePokemon(ctx, CreatePokemonParams{
			Name:       arg.Name,
			Type1:      arg.Type1,
			Type2:      arg.Type2,
			Total:      arg.Total,
			Hp:         arg.Hp,
			Attack:     arg.Attack,
			Defense:    arg.Defense,
			SpAtk:      arg.SpAtk,
			SpDef:      arg.SpDef,
			Speed:      arg.Speed,
			Generation: arg.Generation,
			Legendary:  arg.Legendary,
		})
		if err != nil {
			return err
		}
		return err
	})

	return result, err
}
