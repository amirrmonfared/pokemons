package reader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const path = "../../data/testpokemon.csv"

func TestExporter(t *testing.T) {
	data, err := Exporter(path)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}
