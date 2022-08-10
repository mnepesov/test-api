package config

import (
	"errors"
	"os"
	"strconv"
	
	"github.com/joho/godotenv"
)

var (
	EnvNotFound = errors.New("environment not found")
)

type AppConfig struct {
	HTTP     HTTP
	Postgres Postgres
	NASA     NASA
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  bool
}

type HTTP struct {
	Port string
}

type NASA struct {
	ApiKey string
}

func NewAppConfig(env string) (*AppConfig, error) {
	if err := godotenv.Load(env + ".env"); err != nil {
		return nil, err
	}
	
	var cfg AppConfig
	if len(os.Getenv("PORT")) > 0 {
		cfg.HTTP.Port = os.Getenv("PORT")
	} else {
		cfg.HTTP.Port = "8080"
	}
	
	if len(os.Getenv("POSTGRES_HOST")) > 0 {
		cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	} else {
		return nil, EnvNotFound
	}
	
	if len(os.Getenv("POSTGRES_PORT")) > 0 {
		cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	} else {
		cfg.Postgres.Port = "5432"
	}
	
	if len(os.Getenv("POSTGRES_USERNAME")) > 0 {
		cfg.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
	} else {
		return nil, EnvNotFound
	}
	
	if len(os.Getenv("POSTGRES_PWD")) > 0 {
		cfg.Postgres.Password = os.Getenv("POSTGRES_PWD")
	} else {
		return nil, EnvNotFound
	}
	
	if len(os.Getenv("POSTGRES_DB_NAME")) > 0 {
		cfg.Postgres.DBName = os.Getenv("POSTGRES_DB_NAME")
	} else {
		return nil, EnvNotFound
	}
	
	if len(os.Getenv("POSTGRES_SSL_MODE")) > 0 {
		val, err := strconv.ParseBool(os.Getenv("POSTGRES_SSL_MODE"))
		if err != nil {
			return nil, err
		}
		
		cfg.Postgres.SSLMode = val
	} else {
		cfg.Postgres.SSLMode = false
	}
	
	if len(os.Getenv("NASA_API_KEY")) > 0 {
		cfg.NASA.ApiKey = os.Getenv("NASA_API_KEY")
	} else {
		return nil, EnvNotFound
	}
	
	return &cfg, nil
}
