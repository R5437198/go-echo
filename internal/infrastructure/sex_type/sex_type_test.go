package sex_type

import (
	"github.com/stretchr/testify/assert"
	"go-echo/internal/infrastructure/orm"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	// Golandからテストだと接続ホストが異なるため
	err := os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	assert.NoError(t, err)
	defer os.Clearenv()

	db := orm.New()
	assert.NotEmpty(t, db)

	assert.NoError(t, err)
	assert.Equal(t, &repository{DB: db.DB}, New(db.DB))
}

func Test_repository_FetchAll(t *testing.T) {
	t.Parallel()

	// Golandからテストだと接続ホストが異なるため
	err := os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	assert.NoError(t, err)
	defer os.Clearenv()

	db := orm.New()
	assert.NotEmpty(t, db)

	r := New(db.DB)
	res, err := r.FetchAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
