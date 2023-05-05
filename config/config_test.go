package config

import (
	// "os"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	// t.Run("test config success", func(t *testing.T) {
	// 	// os.Setenv("port", "localhost:5000")
	// 	// os.Setenv("storage", "mysql")
	// 	// defer os.Unsetenv("ENV_VAR")

    // 	// require.Equal(t, "localhost:5000", os.Getenv("port"))
    // 	// require.Equal(t, "mysql", os.Getenv("storage"))

	// 	config, err := LoadConfig(".env", "./../config")
	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, config)
	// })
	t.Run("test config failed 1", func(t *testing.T) {
		config, err := LoadConfig(".env", "./../config")
		require.Error(t, err)
		require.Empty(t, config)
	})

	t.Run("test config failed 2", func(t *testing.T) {
		config, err := LoadConfig(".env", "../config")
		require.Error(t, err)
		require.Empty(t, config)
	})
}