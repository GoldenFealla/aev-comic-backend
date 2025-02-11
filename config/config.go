/*
This config package is mainly for loading environment variable (included env file) and json file

Call

	config.Load()

to load the config
*/
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

const (
	corsPath = "./config/cors.yaml"
	envPath  = "./config/.env"
)

type Environment struct {
	Mode        string
	PostgresURI string
}

type Config struct {
	CORS middleware.CORSConfig
	Env  Environment
}

func loadCORS() middleware.CORSConfig {
	filename, _ := filepath.Abs(corsPath)
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var cors middleware.CORSConfig

	err = yaml.Unmarshal(yamlFile, &cors)
	if err != nil {
		log.Fatalln(err)
	}

	return cors
}

func loadEnv() Environment {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("no env file in ./config/.env. Try to load from root")
		godotenv.Load()
	}

	mode := os.Getenv("MODE")
	postgresURI := os.Getenv("POSTGRES_URI")

	if postgresURI == "" {
		log.Println("no POSTGRES_URI environment provided.")
	}

	return Environment{
		Mode:        mode,
		PostgresURI: postgresURI,
	}
}

func Load() *Config {
	cors := loadCORS()
	env := loadEnv()

	return &Config{
		CORS: cors,
		Env:  env,
	}
}
