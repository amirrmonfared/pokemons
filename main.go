package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/amirrmonfared/pokemons/api"
	db "github.com/amirrmonfared/pokemons/db/sqlc"
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
	log.Println("connected to server on port 8080")
	fmt.Println("--------------------------------------")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	log.Println("connected to database")
	fmt.Println("--------------------------------------")

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		fmt.Println("cannot connect to server", err)
	}

	go reader.Impoerter(conn, path)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	defer conn.Close()
}
