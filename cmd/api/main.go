package main

import (
	"github.com/VusalShahbazov/toimi-api/internal/helpers"
	"github.com/VusalShahbazov/toimi-api/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variable from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	//Init config
	conf := initConfig()

	// New Api server Struct
	srv := server.NewApiServer(conf)

	// Listen and server
	err = srv.Run()

	log.Fatalln(err)

}

func initConfig() *server.Config {
	dbConfig := server.DBConfig{
		Host:     helpers.Get(os.Getenv("DB_HOST"), "localhost"),
		Username: helpers.Get(os.Getenv("DB_USERNAME"), "postgres"),
		Database: helpers.Get(os.Getenv("DB_DATABASE"), "postgres"),
		Password: helpers.Get(os.Getenv("DB_PASSWORD"), ""),
		Port:     helpers.Get(os.Getenv("DB_PORT"), "5432"),
	}

	config := server.NewConfig(
		helpers.Get(os.Getenv("BIND_ADDR"), ":8080"),
		dbConfig,
	)

	return config
}
