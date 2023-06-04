package main

import (
	"airbnb/config"
	db_mysql "airbnb/pkg/db/mysql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Cannot connected to db")
	}

	if db, err := db_mysql.ConnectDB(cfg); err != nil {
		log.Fatalln("Error is: ", err)
	} else {
		log.Println("Connected to db: ", db)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8081")
}
