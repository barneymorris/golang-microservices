package main

import (
	"authentication-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const WEB_PORT = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	log.Printf("Starting authentication service on port %s\n", WEB_PORT)

	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", WEB_PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}