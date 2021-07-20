package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	_, err := ConnectDB()

	if err != nil {
		os.Exit(1)
	}

	r := gin.Default()
	ConfigRoutes(r)
	log.Println(r.Run())
}
