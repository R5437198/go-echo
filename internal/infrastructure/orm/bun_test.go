package orm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	// Golandからテストだと接続ホストが異なるため
	err := os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	assert.NoError(t, err)
	defer os.Clearenv()

	assert.NotEmpty(t, New())
}
