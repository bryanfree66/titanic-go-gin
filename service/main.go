package main

import (
	"github.com/bryanfree66/titanic-go-gin/service/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/passengers/", api.GetPassengers)
	// TODO: HANDLE CONNECTION ERRORS
	r.Run()
}
