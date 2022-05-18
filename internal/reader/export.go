package reader

import (
	"os"

	"github.com/gocarina/gocsv"
)

func Exporter(path string) ([]*Pokemons, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	entries := []*Pokemons{}

	// Load pokemons from file
	if err := gocsv.UnmarshalFile(file, &entries); err != nil {
		panic(err)
	}

	return entries, err
}
