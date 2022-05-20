-- name: CreatePokemon :one
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
) RETURNING *;

-- name: GetPokemon :one
SELECT * FROM pokemons
WHERE id = $1 LIMIT 1;

-- name: GetPokemonByName :one
SELECT * FROM pokemons
WHERE name = $1 LIMIT 1;

-- name: ListPokemons :many
SELECT * FROM pokemons
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListPokemonsByAbility :many
SELECT * FROM pokemons
WHERE hp = $1 OR attack = $2 OR defense = $3
ORDER BY id;

-- name: DeletePokemon :exec
DELETE FROM pokemons
WHERE id = $1;
