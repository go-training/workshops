package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"gin-http/config"
	"gin-http/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Load()
	var server = flag.Bool("server", false, "enable server")
	var ping = flag.Bool("ping", false, "ping server")
	flag.Parse()

	hander := router.Init()

	s := http.Server{
		Addr:         ":" + config.Setting.Server.Port,
		Handler:      hander,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if *server {
		s.ListenAndServe()
	}

	if *ping {
		resp, err := http.Get("http://localhost:" + config.Setting.Server.Port + "/healthz")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatal("can't connect to server")
		}

		log.Println("connected to the server")
	}
}
