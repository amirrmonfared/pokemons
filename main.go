package main

import (
	"database/sql"
	"log"

	"github.com/amirrmonfared/pokemons/internal/reader"
	"github.com/amirrmonfared/pokemons/util"
	_ "github.com/lib/pq"
)

const path = "./data/pokemon.csv"

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	log.Println("connected to DB...")

	reader.Impoerter(conn, path)

	defer conn.Close()
}
