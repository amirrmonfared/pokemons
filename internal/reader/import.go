package reader

import (
	"context"
	"fmt"
	"log"
	"strings"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
)

var PokemonsSl = make([]*Pokemons, 0, 700)

// Importer is importing data to database
func Impoerter(store db.Store, path string) error {
	data, err := reviewer(path)
	if err != nil {
		log.Println(err)
	}

	for _, pk := range data {
		store.ImportPokemon(context.Background(), db.CreatePokemonParams{
			Name:       pk.Name,
			Type1:      pk.Type1,
			Type2:      pk.Type2,
			Total:      pk.Total,
			Hp:         pk.HP,
			Attack:     pk.Attack,
			Defense:    pk.Defense,
			SpAtk:      pk.SpAtk,
			SpDef:      pk.SpDef,
			Speed:      pk.Speed,
			Generation: pk.Generation,
			Legendary:  pk.Legendary,
		})
	}

	fmt.Println("Hey Professor everything is Done!!")

	return nil
}

func reviewer(path string) ([]*Pokemons, error) {
	data, err := Exporter(path)
	if err != nil {
		log.Println(err)
	}

	for _, pk := range data {

		if pk.Legendary == true {
			fmt.Printf("Legend Pokemon %s excluded\n", pk.Name)
		}
		if isTrue(pk.Type1, pk.Type2, "Ghost") == true {
			fmt.Printf("Ghost Pokemon %s excluded\n", pk.Name)
		}
		if isTrue(pk.Type1, pk.Type2, "Steel") == true {
			pk.HP = pk.HP * 2
		}
		if isTrue(pk.Type1, pk.Type2, "Fire") {
			pk.Attack = decreaseByPercentage(pk.Attack)
		}
		if isTrue(pk.Type1, pk.Type2, "Flying") || isTrue(pk.Type1, pk.Type2, "Bug") {
			pk.SpAtk = increaseByPercentage(pk.SpAtk)
		}
		if str := strings.Split(pk.Name, ""); str[0] == "G" {
			outG := strings.ReplaceAll(pk.Name, "G", "")
			outg := strings.ReplaceAll(outG, "g", "")

			for i := 0; i < len(outg); i++ {
				pk.Defense = pk.Defense + 5
			}
		}
		PokemonsSl = append(PokemonsSl, pk)
	}

	return PokemonsSl, nil
}

func isTrue(pokemonType1, pokemonType2 string, expectedType string) bool {
	if pokemonType1 == expectedType || pokemonType2 == expectedType {
		return true
	}
	return false
}

func decreaseByPercentage(a int32) int32 {
	percent := 10 / float64(a)
	real := percent * 100
	sub := float64(a) - real
	return int32(sub)
}

func increaseByPercentage(a int32) int32 {
	percent := 10 / float64(a)
	real := percent * 100
	sub := float64(a) + real
	return int32(sub)
}
