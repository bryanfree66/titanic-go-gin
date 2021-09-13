package handler

import (
	"github.com/bryanfree66/titanic-go-gin/app/model"
	"github.com/bryanfree66/titanic-go-gin/app/model/errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
)

// Handler struct holds required services for handler to function
type Handler struct {
	PassengerService model.PassengerService
}

// Config will hold services that will eventually be injected
type Config struct {
	R                *gin.Engine
	PassengerService model.PassengerService
}

// NewHandler initializes the handler with required injected services
func NewHandler(c *Config) {
	h := &Handler{
		PassengerService: c.PassengerService,
	}

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
	uuidString := c.Param("uuid")

	uuid, err := primitive.ObjectIDFromHex(uuidString)
	if err != nil {
		log.Printf("Unable to get object id from url: %v\n%v", uuid, err)
		e := errors.NewInternal()

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	p, err := h.PassengerService.Get(c, uuid)
	if err != nil {
		log.Printf("Unable to find passenger: %v\n%v", uuid, err)
		e := errors.NewNotFound("passenger", uuid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"passenger": p,
	})
}
