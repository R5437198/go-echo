package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Parallel()
	t.Run(".envのREST_CONTAINER_PORT,DB_URLを返却する", func(t *testing.T) {
		e, err := Load()
		assert.NoError(t, err)

		assert.Equal(
			t,
			&Env{
				ServerPORT: "80",
				DbURL:      "postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable",
			},
			e,
		)
	})
}

func TestENV_DbUrl(t *testing.T) {
	t.Parallel()
	t.Run(".envのDB_URLが返却される", func(t *testing.T) {
		e, err := Load()
		assert.NoError(t, err)

		assert.Equal(
			t,
			"postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable",
			e.DbUrl(),
		)
	})
}

func TestENV_ServerPort(t *testing.T) {
	t.Parallel()
	t.Run(".envのREST_CONTAINER_PORTが返却される", func(t *testing.T) {
		e, err := Load()
		assert.NoError(t, err)
		serverPort, err := e.ServerPort()
		assert.NoError(t, err)

		assert.Equal(
			t,
			80,
			serverPort,
		)
	})
}
