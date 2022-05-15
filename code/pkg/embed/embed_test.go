package embed

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed static
var embedFS embed.FS

func TestEmbed(t *testing.T) {
	entities, err := embedFS.ReadDir("static")
	assert.NoError(t, err)

	for _, entity := range entities {
		info, err := embedFS.ReadFile("static/" + entity.Name())
		assert.NoError(t, err)
		t.Logf("file name: %s, values: %s\n", entity.Name(), info)
	}
}
