package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME"`
	Port        int32  `env:"PORT" envDefault:"3350"`
	Database    struct {
		Host     string `env:"MYSQL_HOST,required"`
		Port     int32  `env:"MYSQL_PORT,required"`
		Database string `env:"MYSQL_DATABASE_NAME,required"`
		UserName string `env:"MYSQL_USER_NAME,required"`
		Password string `env:"MYSQL_PASSWORD,required"`
	}
}

func (c *Config) New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if err := env.Parse(c); err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
		return nil, err
	}

	return c, nil
}
