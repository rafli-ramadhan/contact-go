package client

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("test db client success", func(t *testing.T) {
		config, err := GetDB("mysql").GetMysqlConnection()
		require.NoError(t ,err)
		require.NotEmpty(t, config)
	})

	t.Run("test db client failed", func(t *testing.T) {
		config, err := GetDB("mysqld").GetMysqlConnection()
		require.Error(t ,err)
		require.Empty(t, config)
	})
}