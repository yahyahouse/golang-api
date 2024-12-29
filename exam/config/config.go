package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
	"time"
)

func init() {
	viper.SetConfigName("default")
	viper.SetConfigType("env")
	viper.AddConfigPath("C:\\Users\\phinc\\OneDrive\\Documents\\Phincon\\LATIHAN\\latihan-go\\exam\\config\\")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file", err)
	}
	viper.AutomaticEnv()

}

type DBConfig struct {
	User, Password          string
	ConString               string
	TableName               string
	MaxOpenCon              int
	MaxIdleCon              int
	ConnIdleTimeoutDuration time.Duration
	QueryTimeoutDuration    time.Duration
}
type KafkaConfig struct {
	ServerAddress string
	GroupId       string
}
type Config struct {
	DBConfig        DBConfig
	KafkaConfig     KafkaConfig
	RedisConfig     RedisConfig
	ServerCfg       ServerConfig
	Port            string
	IsByPassDBRedis bool
	IsValid         bool
}
type ServerConfig struct {
	ReadTimeout, WriteTimeout string
}

func GetConfig() Config {
	return Config{
		DBConfig:        getDBConfig(),
		KafkaConfig:     GetKafkaConfig(),
		RedisConfig:     getRedisConfig(),
		ServerCfg:       getServerConfig(),
		Port:            viper.GetString("port"),
		IsByPassDBRedis: viper.GetBool("is_by_pass_db_redis"),
		IsValid:         true,
	}
}
func getDBConfig() DBConfig {
	return DBConfig{
		User:                    viper.GetString("database_user"),
		Password:                viper.GetString("database_password"),
		ConString:               viper.GetString("database_url"),
		MaxOpenCon:              viper.GetInt("database_max_open_connections"),
		MaxIdleCon:              viper.GetInt("database_max_idle_connections"),
		QueryTimeoutDuration:    viper.GetDuration("database_set_timeout"),
		ConnIdleTimeoutDuration: viper.GetDuration("database_conn_idle_timeout_duration"),
	}
}
func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		ServerAddress: viper.GetString("kafka_server_address"),
		GroupId:       viper.GetString("kafka_group_id"),
	}
}

type RedisConfig struct {
	MasterName             string
	SentinelAddrs          []string
	User, Password         string
	Database               int
	PoolSize, MaxIdleConns int
	Timeout                time.Duration
	TimeToLive             time.Duration
}

func getRedisConfig() RedisConfig {
	redisDB, errCfg := strconv.Atoi(viper.GetString("REDIS_DB"))
	if errCfg != nil {
		redisDB = 0
	}

	poolSize, errCfg := strconv.Atoi(viper.GetString("REDIS_POOL_SIZE"))
	if errCfg != nil {
		poolSize = 100
	}

	maxIdleConns, errCfg := strconv.Atoi(viper.GetString("REDIS_MAX_IDLE_CONNECTIONS"))
	if errCfg != nil {
		poolSize = 10
	}

	timeout, errCfg := time.ParseDuration(viper.GetString("REDIS_TIMEOUT_DURATION"))
	if errCfg != nil {
		panic(fmt.Sprintf("Failed load conf REDIS_TIMEOUT_DURATION %s", errCfg))
	}

	user := viper.GetString("REDIS_USER")

	password := viper.GetString("REDIS_PASSWORD")

	sentinalAddr := viper.GetString("REDIS_SENTINEL_ADDR")
	sentinalAddrs := strings.Split(sentinalAddr, ",")

	masterName := viper.GetString("REDIS_MASTER_NAME")

	return RedisConfig{
		User:          user,
		MasterName:    masterName,
		SentinelAddrs: sentinalAddrs,
		Password:      password,
		Database:      redisDB,
		Timeout:       timeout,
		PoolSize:      poolSize,
		MaxIdleConns:  maxIdleConns,
	}

}
func getServerConfig() ServerConfig {
	return ServerConfig{
		ReadTimeout:  viper.GetString("server_read_timeout"),
		WriteTimeout: viper.GetString("server_write_timeout"),
	}
}
