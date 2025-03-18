package main

import (
	httpapi "key-haven-back/internal/http"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := fx.New(
		httpapi.Module,
	)

	app.Run()
}
