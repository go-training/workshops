package main

import (
	"net/http"

	"gin-http-server/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	// "github.com/kelseyhightower/envconfig"
)

// type (
// 	// Config provides the system configuration.
// 	Server struct {
// 		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
// 	}
// )

func main() {
	// server := &Server{}
	// envconfig.MustProcess("", server)
	// config.Load()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Setting.Server.Port) // listen on 0.0.0.0:8088
}
