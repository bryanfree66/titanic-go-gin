package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// PassengerService defines methods the handler layer expects
type PassengerService interface {
	GetAll() ([]Passenger, error)
	Get(uuid primitive.ObjectID) (*Passenger, error)
}

// PassengerRepository defines methods the service layer expects
type PassengerRepository interface {
	GetAll() ([]Passenger, error)
	FindByID(uid uuid.UUID) (*Passenger, error)
}
