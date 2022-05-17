package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/amirrmonfared/pokemons/util"
	"github.com/gocarina/gocsv"
	_ "github.com/lib/pq"
)

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

	Entries()

	defer conn.Close()
}

type Pokemons struct {
	Id         int32  `csv:"#"`
	Name       string `csv:"Name"`
	Type1      string `csv:"Type 1"`
	Type2      string `csv:"Type 2"`
	Total      int32  `csv:"Total"`
	HP         int32  `csv:"HP"`
	Attack     int32  `csv:"Attack"`
	Defense    int32  `csv:"Defense"`
	SpAtk      int32  `csv:"Sp. Atk"`
	SpDef      int32  `csv:"Sp. Def"`
	Speed      int32  `csv:"Speed"`
	Generation int32  `csv:"Generation"`
	Legendary  bool   `csv:"Legendary"`
}


func Entries() {
	file, err := os.OpenFile("./data/pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	entries := []*Pokemons{}

	// Load pokemons from file
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		panic(err)
	}

	for _, pokemon := range entries {
		fmt.Println("Hello", pokemon.Name)
		
	}

	// Go to the start of the file
	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	// Get all clients as CSV string
	csvContent, err := gocsv.MarshalString(&entries)
	if err != nil {
		panic(err)
	}

	fmt.Println(csvContent)
}