package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StrinDataBaseConnection = ""
	Porta                   = 0
)

func Load() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Porta, error = strconv.Atoi(os.Getenv("API_PORT"))
	// TODO - PUT RADOM PORT
	if error != nil {
		Porta = 9000
	}

	StrinDataBaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
