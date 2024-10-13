package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sakthi-lucia0567/animart/internal/config"
	"github.com/sakthi-lucia0567/animart/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := server.New()

	fmt.Printf("Server starting on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
