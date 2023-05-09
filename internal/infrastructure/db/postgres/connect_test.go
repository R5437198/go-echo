package postgres

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("DBと接続が可能", func(t *testing.T) {
		// Golandからテストだと接続ホストが異なるため
		err := os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
		assert.NoError(t, err)
		defer os.Clearenv()

		db, err := New()
		defer db.Close()
		assert.NoError(t, err)
		assert.NotEmpty(t, db)
	})
}
