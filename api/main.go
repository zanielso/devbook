package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println(config.StrinDataBaseConnection)

	fmt.Println(config.Porta)

	fmt.Println("Running API!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
