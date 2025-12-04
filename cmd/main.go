package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "info: ", log.LstdFlags)
	srv := server.NewServer(logger)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
