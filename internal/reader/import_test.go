package reader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReviewer(t *testing.T) {
	data, err := reviewer(path)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestImporter(t *testing.T) {
	err := Impoerter(testStore, path)
	require.NoError(t, err)
}
