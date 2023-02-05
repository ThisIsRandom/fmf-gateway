package main

import (
	"log"

	"github.com/thisisrandom/fmf-gateway/clients"
)

func main() {
	userClient, err := clients.InitTokenClient(":8080")

	if err != nil {
		log.Fatalf("err: %s", err)
	}

	res, err := userClient.GenerateToken("test", "test")

	if err != nil {
		log.Fatalf("err: %s", err)
	}

	log.Printf("token: %s", res.Token)
}
