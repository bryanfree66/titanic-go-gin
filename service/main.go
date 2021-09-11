package main

import (
	"github.com/bryanfree66/titanic-go-gin/service/api"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/api/passengers/", api.GetPassengers)
	runErr := r.Run()
	if runErr != nil {
		log.Fatalf("There was a problem starting the server: %v", runErr)
	}
}
