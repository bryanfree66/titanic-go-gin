package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PassengerService defines methods the handler layer expects
type PassengerService interface {
	// GetAll(ctx context.Context) ([]Passenger, error)
	Get(ctx context.Context, uuid primitive.ObjectID) (*Passenger, error)
}

// PassengerRepository defines methods the service layer expects
type PassengerRepository interface {
	// GetAll(ctx context.Context) ([]Passenger, error)
	FindByID(ctx context.Context, uuid primitive.ObjectID) (*Passenger, error)
}
