package reader

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
	"github.com/amirrmonfared/pokemons/util"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB
var testStore db.Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
