package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

func init() {
	secretKey := make([]byte, 64)

	if _, err := rand.Read(secretKey); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
