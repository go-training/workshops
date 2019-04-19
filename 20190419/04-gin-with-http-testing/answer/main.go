package main

import (
	"log"
	"net/http"
	"time"

	"gin-http-server/config"
	"gin-http-server/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := &http.Server{
		Addr:    ":" + config.Setting.Server.Port,
		Handler: router.Handler(),
		// for upload data
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
