package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestViperReadsConfigFileSuccessfully(t *testing.T) {
	viper.SetConfigName("default")
	viper.SetConfigType("env")
	viper.AddConfigPath("C:\\Users\\phinc\\OneDrive\\Documents\\Phincon\\LATIHAN\\latihan-go\\exam\\config\\")

	err := viper.ReadInConfig()

	assert.NoError(t, err, "Expected no error while reading the config file")
}

func TestViperConfigFileNotFound(t *testing.T) {
	viper.SetConfigName("nonexistent")
	viper.SetConfigType("env")
	viper.AddConfigPath("C:\\Users\\phinc\\OneDrive\\Documents\\Phincon\\LATIHAN\\latihan-go\\exam\\config\\")

	err := viper.ReadInConfig()

	assert.Error(t, err, "Expected an error while reading a nonexistent config file")
}

func TestGetConfigRetrievesDBConfig(t *testing.T) {
	viper.Set("database_user", "testuser")
	viper.Set("database_password", "testpassword")
	viper.Set("database_url", "localhost:5432/testdb")
	viper.Set("database_max_open_connections", 10)
	viper.Set("database_max_idle_connections", 5)
	viper.Set("database_set_timeout", "30s")
	viper.Set("database_conn_idle_timeout_duration", "10m")

	cfg := GetConfig()

	assert.NotNil(t, cfg.DBConfig)
	assert.Equal(t, "testuser", cfg.DBConfig.User)
	assert.Equal(t, "testpassword", cfg.DBConfig.Password)
	assert.Equal(t, "localhost:5432/testdb", cfg.DBConfig.ConString)
	assert.Equal(t, 10, cfg.DBConfig.MaxOpenConns)
	assert.Equal(t, 5, cfg.DBConfig.MaxIdleConns)
	assert.Equal(t, 30*time.Second, cfg.DBConfig.Timeout)
	assert.Equal(t, 10*time.Minute, cfg.DBConfig.ConnIdleTimeoutDuration)
}
func TestGetConfigFailsToRetrievePort(t *testing.T) {
	viper.Reset()

	cfg := GetConfig()

	assert.Empty(t, cfg.Port)
}
