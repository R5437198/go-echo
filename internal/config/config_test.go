package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_DbUrl(t *testing.T) {
	t.Parallel()
	t.Run("ConfigのDbUrlを返却する", func(t *testing.T) {
		c, err := New()
		assert.NoError(t, err)
		assert.Equal(t,
			"postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable",
			c.DbUrl(),
		)
	})
}

func TestConfig_Driver(t *testing.T) {
	t.Parallel()
	t.Run("ConfigのDriverを返却する", func(t *testing.T) {
		c, err := New()
		assert.NoError(t, err)
		assert.Equal(t, "postgres", c.Driver())
	})
}

func TestConfig_Port(t *testing.T) {
	t.Parallel()
	t.Run("ConfigのPortを返却する", func(t *testing.T) {
		c, err := New()
		assert.NoError(t, err)
		assert.Equal(t, ":80", c.Port())
	})
}

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("Configを返却する", func(t *testing.T) {
		c, err := New()
		assert.NoError(t, err)
		assert.Equal(t, &Config{
			Server: Server{
				Port: 80,
			},
			DB: DB{
				Driver: "postgres",
				Url:    "postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable",
			},
		}, c)
	})
}
