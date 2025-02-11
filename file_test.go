package dadjokes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFileJokerDefault(t *testing.T) {
	r := require.New(t)
	fj, err := NewFileJokerDefault()
	r.NoError(err)
	t.Logf("joke: %s", fj.GetJoke())
	r.NotEmpty(fj.GetJoke())
}
