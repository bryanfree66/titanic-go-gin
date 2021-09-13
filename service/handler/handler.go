package handler

import (
	"context"
	"fmt"
	"github.com/bryanfree66/titanic-go-gin/service/helpers"
	"github.com/bryanfree66/titanic-go-gin/service/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

// database and collection constants
const (
	database   = "titanic"
	collection = "passengers"
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
	g := c.R.Group("/api/passengers")

	// route handlers
	g.GET("/", h.Passengers)
}

// Passengers handler returns all passengers
func (h *Handler) Passengers(c *gin.Context) {
	log.Println("Handling GET request for all passengers.")

	// Get the database client
	client, ctx, cancel := helpers.GetConnection()
	defer cancel()
	defer func() {
		log.Println("Disconnecting Mongo client.")
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("Error disconnecting MongoDB client")
			panic(err)
		}
	}()

	// Get the data from MongoDB client
	cur, getErr := client.Database(database).Collection(collection).Find(context.TODO(), bson.M{})
	if getErr != nil {
		msg := fmt.Sprintf("Error retrieving passengers.")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	// Decode BSON
	var passengers []model.Passenger
	for cur.Next(context.TODO()) {
		var passenger model.Passenger
		decodeErr := cur.Decode(&passenger)
		if decodeErr != nil {
			log.Printf(" Error decoding passenger: %v", decodeErr)
			c.IndentedJSON(http.StatusInternalServerError, decodeErr)
		}
		passengers = append(passengers, passenger)
	}

	// Return successful result
	log.Println("Successfully handled GET request for all passengers")
	c.IndentedJSON(http.StatusOK, passengers)
}
