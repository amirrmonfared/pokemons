package reader

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
)

var PokemonsSl = make([]*Pokemons, 0, 700)

// Importer is importing data to database
func Impoerter(conn *sql.DB, path string) error {
	store := db.NewStore(conn)
	data, err := reviewer(path)
	if err != nil {
		log.Println(err)
	}

	for _, a := range data {
		store.ImportPokemon(context.Background(), db.CreatePokemonParams{
			Name:       a.Name,
			Type1:      a.Type1,
			Type2:      a.Type2,
			Total:      a.Total,
			Hp:         a.HP,
			Attack:     a.Attack,
			Defense:    a.Defense,
			SpAtk:      a.SpAtk,
			SpDef:      a.SpDef,
			Speed:      a.Speed,
			Generation: a.Generation,
			Legendary:  a.Legendary,
		})
	}

	fmt.Println("Hey Professor everything is Done!!")

	return nil
}

// reviewer checking if rules
func reviewer(path string) ([]*Pokemons, error) {
	data, err := Exporter(path)
	if err != nil {
		log.Println(err)
	}

	for _, a := range data {
		str := strings.Split(a.Name, "")

		if a.Legendary == true {
			//Exclude Legendary Pokémon
			fmt.Printf("Legend Pokemon %s excluded\n", a.Name)

		} else if a.Type1 == "Ghost" || a.Type2 == "Ghost" {
			//Exclude Pokémon of Type: Ghost
			fmt.Printf("Ghost Pokemon %s excluded\n", a.Name)

		} else if a.Type1 == "Steel" || a.Type2 == "Steel" {
			//For Pokémon of Type: Steel, double their HP
			a.HP = a.HP * 2
			PokemonsSl = append(PokemonsSl, a)
			
		} else if a.Type1 == "Fire" || a.Type2 == "Fire" {
			//For Pokémon of Type: Fire, lower their Attack by 10%
			percent := 10 / float64(a.Attack)
			real := percent * 100
			sub := float64(a.Attack) - real
			a.Attack = int32(sub)
			PokemonsSl = append(PokemonsSl, a)

		} else if a.Type2 == "Flying" || a.Type1 == "Flying" {
			//For Pokémon of Type: Flying, increase their Attack Speed by 10%
			percent := 10 / float64(a.SpAtk)
			real := percent * 100
			sum := float64(a.SpAtk) + real
			a.SpAtk = int32(sum)
			PokemonsSl = append(PokemonsSl, a)

		} else if a.Type2 == "Bug" || a.Type1 == "Bug" {
			//For Pokémon of Type: Bug, increase their Attack Speed by 10%
			percent := 10 / float64(a.SpAtk)
			real := percent * 100
			sum := float64(a.SpAtk) + real
			a.SpAtk = int32(sum)
			PokemonsSl = append(PokemonsSl, a)

		} else if str[0] == "G" {
			//For Pokémon that start with the letter **G**, add +5 Defense for every letter in their name (excluding **G**)
			outG := strings.ReplaceAll(a.Name, "G", "")
			outg := strings.ReplaceAll(outG, "g", "")

			for i := 0; i < len(outg); i++ {
				a.Defense = a.Defense + 5
			}
			PokemonsSl = append(PokemonsSl, a)

		} else {

			PokemonsSl = append(PokemonsSl, a)
		}
	}

	return PokemonsSl, nil
}
