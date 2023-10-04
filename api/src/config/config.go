package config

import (
	"fmt"
	"log"
	"math/rand"
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
	if error != nil {
		min := 8000
		max := 9990
		Porta = rand.Intn(max-min) + min
	}

	StrinDataBaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
