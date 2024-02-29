package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Env      string         `yaml:"env" env-default:"local"`
		Postgres PostgresConfig `yaml:"postgres"`
		HTTP     HTTPConfig     `yaml:"http"`
	}

	PostgresConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"` //TODO : read pwd from env
	}

	HTTPConfig struct {
		Host        string        `yaml:"host"`
		Port        string        `yaml:"port"`
		Timeout     time.Duration `yaml:"timeout"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
	}
)

// TODO : REFACTOR CODE BELOW (FOLD CAN BE EMPTY BUT ENV IS NOT)
func MustLoad(folder, env string) *Config {
	if folder == "" || env == "" {
		log.Fatal("FOLDER or ENV is not set")
	}

	configPath := folder + env + ".yaml"

	//check if file exists
	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s\n", configPath)
		} else {
			log.Fatalf("file check error: %v\n", err)
		}
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
