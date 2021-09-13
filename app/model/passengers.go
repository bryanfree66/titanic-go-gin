package model

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SexType string

const (
	Male   SexType = "male"
	Female         = "female"
)

func (st SexType) IsValid() error {
	switch st {
	case Male, Female:
		return nil
	}
	return errors.New("invalid sex")
}

type EmbarkType string

const (
	EmbarcationUnknown EmbarkType = "Unknown"
	Cherbourg                     = "Cherbourg"
	Southampton                   = "Southampton"
	Queenstown                    = "Queenstown"
)

func (et EmbarkType) IsValid() error {
	switch et {
	case EmbarcationUnknown, Cherbourg, Southampton, Queenstown:
		return nil
	}
	return errors.New("invalid embarcation")
}

type PassengerClassType int

const (
	ClassNotKnown PassengerClassType = iota
	First
	Second
	Third
)

func (pc PassengerClassType) IsValid() error {
	switch pc {
	case ClassNotKnown, First, Second, Third:
		return nil
	}
	return errors.New("invalid passenger class")
}

type DeckType string

const (
	DeckNotKnown DeckType = "Unknown"
	A                     = "A"
	B                     = "B"
	C                     = "C"
	D                     = "D"
	E                     = "E"
	F                     = "F"
	G                     = "G"
)

func (dt DeckType) IsValid() error {
	switch dt {
	case DeckNotKnown, A, B, C, D, E, F, G:
		return nil
	}
	return errors.New("invalid deck")
}

type TripData struct {
	PassengerClass PassengerClassType `json:"passengerClass,omitempty" bson:"PassengerClass,omitempty"`
	Deck           DeckType           `json:"deck,omitempty" bson:"Deck,omitempty"`
	Cabin          []int              `json:"cabin,omitempty" bson:"Cabin,omitempty"`
	GroupSize      int                `json:"groupSize,omitempty" bson:"GroupSize,omitempty"`
	Ticket         string             `json:"ticket,omitempty" bson:"Ticket,omitempty"`
	Fare           float32            `json:"fare,omitempty" bson:"Fare,truncate"`
	Embarked       EmbarkType         `json:"embarked,omitempty" bson:"Embarked,omitempty"`
}

type Passenger struct {
	UUID        primitive.ObjectID `json:"uuid,omitempty" bson:"_id,omitempty"`
	PassengerId int                `json:"passengerId,omitempty" bson:"PassengerId,omitempty"`
	Survived    bool               `json:"survived,omitempty" bson:"Survived,omitempty"`
	Title       string             `json:"title,omitempty" bson:"Title,omitempty"`
	FirstName   string             `json:"firstName,omitempty" bson:"FirstName,omitempty"`
	MiddleName  string             `json:"middleNameName,omitempty" bson:"MiddleNameName,omitempty"`
	LastName    string             `json:"lastName,omitempty" bson:"LastName,omitempty"`
	Sex         SexType            `json:"sex,omitempty"  bson:"Sex,omitempty"`
	Age         int                `json:"age,omitempty"  bson:"Age,truncate"`
	TripData    TripData           `json:"tripData"  bson:"TripData"`
}
