package client

import (
	"log"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("test db client success", func(t *testing.T) {
		config, err := GetDBConnection("mysql").GetMysqlConnection()
		require.NoError(t ,err)
		require.NotEmpty(t, config)
	})

	t.Run("test db client failed", func(t *testing.T) {
		config, err := GetDBConnection("").GetMysqlConnection()
		require.Error(t, err)
		require.Empty(t, config)
	})

	t.Run("test db gorm client success", func(t *testing.T) {
		config, err := GetDBConnection("mysql-gorm").GetMysqlGormConnection()
		require.NoError(t ,err)
		require.NotEmpty(t, config)
	})

	t.Run("test db gorm client failed", func(t *testing.T) {
		_, err := GetDBConnection("").GetMysqlGormConnection()
		log.Println(err)
		require.Error(t ,err)
	})
}