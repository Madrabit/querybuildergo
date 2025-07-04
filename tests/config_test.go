package tests

import (
	"github.com/stretchr/testify/assert"
	"querybuilder/internal/common"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Run("empty .env - full ram env", func(t *testing.T) {
		t.Setenv("DB_DATABASE", "ibd_renew1")
		t.Setenv("DB_SERVER", "10.0.0.1")
		t.Setenv("DB_PORT", "1433")
		t.Setenv("DB_USER", "sa")
		t.Setenv("DB_PASS", "hf,jnfflvbyf")
		t.Setenv("SERVER_ADDRESS", "localhost")
		t.Setenv("SERVER_PORT", "8080")
		cnf, err := common.Load()
		assert.NoError(t, err)
		assert.Equal(t, "ibd_renew1", cnf.DB.Database)
	})
	t.Run("empty .env and empty environment", func(t *testing.T) {
		t.Setenv("DB_DATABASE", "")
		t.Setenv("DB_SERVER", "")
		t.Setenv("DB_PORT", "0")
		t.Setenv("DB_USER", "")
		t.Setenv("DB_PASS", "")
		t.Setenv("DB_PASS", "")
		t.Setenv("SERVER_ADDRESS", "")
		t.Setenv("SERVER_PORT", "0")
		cnf, err := common.Load()
		assert.NoError(t, err)
		assert.Equal(t, common.Config{}, cnf)
		assert.Empty(t, cnf.DB.Database)
	})
}
