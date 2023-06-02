package complexity_test

import (
	"testing"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/Warashi/compgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFail(t *testing.T) {
	cfg, err := config.LoadConfigFromDefaultLocations()
	require.NoError(t, err)

	err = api.Generate(cfg, api.AddPlugin(compgen.New()))
	assert.ErrorIs(t, err, compgen.ErrMulFieldIsNotExist)
}
