package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Handler struct holds required services for handler to function
type Handler struct{}

// Config will hold services that will eventually be injected
type Config struct {
	R *gin.Engine
}

// NewHandler initializes the handler with required injected services
func NewHandler(c *Config) {
	h := &Handler{}

	// Create a group, or base url for all routes
	g := c.R.Group(os.Getenv("PASSENGER_API_URL"))

	// route handlers
	g.GET("/", h.Passengers)
	g.GET("/:uuid", h.Passenger)
}

// Passengers handler returns all passengers
func (h *Handler) Passengers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"handler": "passengers",
	})
}

// Passenger handler returns all passengers
func (h *Handler) Passenger(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(http.StatusOK, gin.H{
		"handler": "passenger: " + uuid,
	})
}
