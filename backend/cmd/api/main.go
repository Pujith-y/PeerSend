package main

import (
	"log"

	"github.com/yenugantirahul/peersend/backend/internal/app"
)

func main() {
	server := app.New()

	log.Println("🚀 PeerSend API starting on :8080")

	server.Logger.Fatal(server.Start(":8080"))
}
