package api

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

const (
	database   = "titanic"
	collection = "passengers"
)

// GetPassengers responds with the list of all albums as JSON.
func GetPassengers(c *gin.Context) {
	// Get the database client
	client, ctx, cancel := helpers.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var passengers []model.Passenger
	cur, getErr := client.Database(database).Collection(collection).Find(context.TODO(), bson.M{})
	if getErr != nil {
		msg := fmt.Sprintf("Error retrieving passengers.")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	// TODO: fix this
	for cur.Next(context.TODO()) {
		var passenger model.Passenger
		decodeErr := cur.Decode(&passenger)
		if decodeErr != nil {
			log.Printf(" Error decoding passenger: %v", decodeErr)
		}
		passengers = append(passengers, passenger)
	}

	c.IndentedJSON(http.StatusOK, passengers)
}
