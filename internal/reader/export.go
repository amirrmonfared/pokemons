package reader

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// Exporter exporting the data from csv file
func Exporter(path string) ([]*Pokemons, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}

	entries := []*Pokemons{}

	// Load pokemons from file
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		log.Panic()
	}

	return entries, nil
}
