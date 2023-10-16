package config

import (
	"os"
	"time"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	Environment      string
	LogLevel         string
	HttpPort         string
	CtxTimeout       int64
	TokenTime        time.Duration
	HeaderKey        string
	PayloadKey       string
	SecretKey        string

}

func Load() Config {
	c := Config{}

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "otajonov"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "clinic"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "quvonchbek"))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "developer"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8081"))
	c.TokenTime = cast.ToDuration(getOrReturnDefault("TOKEN_TIME", time.Duration(6200) * time.Second))
	c.SecretKey = cast.ToString(getOrReturnDefault("SECRET_KEY", "clinic"))
	c.HeaderKey = cast.ToString(getOrReturnDefault("HEADER_KEY", "Authorization"))
	c.PayloadKey = cast.ToString(getOrReturnDefault("PAYLOAD_KEY", ""))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}


