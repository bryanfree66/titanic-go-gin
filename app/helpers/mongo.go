package helpers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	// Timeout in seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s:%s"
)

// GetConnection creates a MongoDB client
func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("DB_USER")
	if len(username) == 0 {
		log.Fatalf("Database username not set.")
	}
	password := os.Getenv("DB_PASSWORD")
	if len(password) == 0 {
		log.Fatalf("Database username not set.")
	}
	endpoint := os.Getenv("DB_ENDPOINT")
	if len(endpoint) == 0 {
		log.Println("DB endpoint set to default: localhost")
		endpoint = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if len(port) == 0 {
		log.Println("DB port set to default: 27017")
		endpoint = "27017"
	}

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, endpoint, port)
	log.Printf("Connection URI: %v", connectionURI)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	// Force a connection to verify the connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to the database!")
	return client, ctx, cancel
}
