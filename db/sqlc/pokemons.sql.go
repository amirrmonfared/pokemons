// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: pokemons.sql

package db

import (
	"context"
)

const createPokemon = `-- name: CreatePokemon :one
INSERT INTO pokemons (
  name,
  type1,
  type2,
  total,
  hp,
  attack,
  defense,
  sp_atk,
  sp_def,
  speed,
  generation,
  legendary  
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id, name, type1, type2, total, hp, attack, defense, sp_atk, sp_def, speed, generation, legendary, created_at
`

type CreatePokemonParams struct {
	Name       string `json:"name"`
	Type1      string `json:"type1"`
	Type2      string `json:"type2"`
	Total      int32  `json:"total"`
	Hp         int32  `json:"hp"`
	Attack     int32  `json:"attack"`
	Defense    int32  `json:"defense"`
	SpAtk      int32  `json:"sp_atk"`
	SpDef      int32  `json:"sp_def"`
	Speed      int32  `json:"speed"`
	Generation int32  `json:"generation"`
	Legendary  bool   `json:"legendary"`
}

func (q *Queries) CreatePokemon(ctx context.Context, arg CreatePokemonParams) (Pokemon, error) {
	row := q.db.QueryRowContext(ctx, createPokemon,
		arg.Name,
		arg.Type1,
		arg.Type2,
		arg.Total,
		arg.Hp,
		arg.Attack,
		arg.Defense,
		arg.SpAtk,
		arg.SpDef,
		arg.Speed,
		arg.Generation,
		arg.Legendary,
	)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type1,
		&i.Type2,
		&i.Total,
		&i.Hp,
		&i.Attack,
		&i.Defense,
		&i.SpAtk,
		&i.SpDef,
		&i.Speed,
		&i.Generation,
		&i.Legendary,
		&i.CreatedAt,
	)
	return i, err
}

const deletePokemon = `-- name: DeletePokemon :exec
DELETE FROM pokemons
WHERE id = $1
`

func (q *Queries) DeletePokemon(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePokemon, id)
	return err
}

const getPokemon = `-- name: GetPokemon :one
SELECT id, name, type1, type2, total, hp, attack, defense, sp_atk, sp_def, speed, generation, legendary, created_at FROM pokemons
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPokemon(ctx context.Context, id int64) (Pokemon, error) {
	row := q.db.QueryRowContext(ctx, getPokemon, id)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type1,
		&i.Type2,
		&i.Total,
		&i.Hp,
		&i.Attack,
		&i.Defense,
		&i.SpAtk,
		&i.SpDef,
		&i.Speed,
		&i.Generation,
		&i.Legendary,
		&i.CreatedAt,
	)
	return i, err
}

const getPokemonByName = `-- name: GetPokemonByName :one
SELECT id, name, type1, type2, total, hp, attack, defense, sp_atk, sp_def, speed, generation, legendary, created_at FROM pokemons
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPokemonByName(ctx context.Context, name string) (Pokemon, error) {
	row := q.db.QueryRowContext(ctx, getPokemonByName, name)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type1,
		&i.Type2,
		&i.Total,
		&i.Hp,
		&i.Attack,
		&i.Defense,
		&i.SpAtk,
		&i.SpDef,
		&i.Speed,
		&i.Generation,
		&i.Legendary,
		&i.CreatedAt,
	)
	return i, err
}

const listPokemons = `-- name: ListPokemons :many
SELECT id, name, type1, type2, total, hp, attack, defense, sp_atk, sp_def, speed, generation, legendary, created_at FROM pokemons
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPokemonsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPokemons(ctx context.Context, arg ListPokemonsParams) ([]Pokemon, error) {
	rows, err := q.db.QueryContext(ctx, listPokemons, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pokemon{}
	for rows.Next() {
		var i Pokemon
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type1,
			&i.Type2,
			&i.Total,
			&i.Hp,
			&i.Attack,
			&i.Defense,
			&i.SpAtk,
			&i.SpDef,
			&i.Speed,
			&i.Generation,
			&i.Legendary,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPokemonsByAbility = `-- name: ListPokemonsByAbility :many
SELECT id, name, type1, type2, total, hp, attack, defense, sp_atk, sp_def, speed, generation, legendary, created_at FROM pokemons
WHERE hp = $1 OR attack = $2 OR defense = $3

ORDER BY id
`

type ListPokemonsByAbilityParams struct {
	Hp      int32 `json:"hp"`
	Attack  int32 `json:"attack"`
	Defense int32 `json:"defense"`
}

func (q *Queries) ListPokemonsByAbility(ctx context.Context, arg ListPokemonsByAbilityParams) ([]Pokemon, error) {
	rows, err := q.db.QueryContext(ctx, listPokemonsByAbility, arg.Hp, arg.Attack, arg.Defense)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pokemon{}
	for rows.Next() {
		var i Pokemon
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type1,
			&i.Type2,
			&i.Total,
			&i.Hp,
			&i.Attack,
			&i.Defense,
			&i.SpAtk,
			&i.SpDef,
			&i.Speed,
			&i.Generation,
			&i.Legendary,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
