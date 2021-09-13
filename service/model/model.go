package model

type Sex int

const (
	Male Sex = iota
	Female
)

type EmbarkCity int

const (
	NoEmbarkedListed EmbarkCity = iota
	Cherbourg
	Southampton
	Queenstown
)

type PassengerClass int

const (
	NoClassListed PassengerClass = iota
	First
	Second
	Third
)

type Deck int

const (
	NoDeckListed Deck = iota
	A
	B
	C
	D
	E
	F
	G
)

type TripData struct {
	PassengerClass PassengerClass `json:"passengerClass,omitempty" bson:"PassengerClass,omitempty"`
	Deck           int            `json:"deck,omitempty" bson:"Deck,omitempty"`
	Cabin          []int          `json:"cabin,omitempty" bson:"Cabin,omitempty"`
	GroupSize      int            `json:"groupSize,omitempty" bson:"GroupSize,omitempty"`
	Ticket         string         `json:"ticket,omitempty" bson:"Ticket,omitempty"`
	Fare           float32        `json:"fare,omitempty" bson:"Fare,truncate"`
	Embarked       EmbarkCity     `json:"embarked,omitempty" bson:"Embarked,omitempty"`
}

type Passenger struct {
	Id          string   `json:"id,omitempty" bson:"_id,omitempty"`
	PassengerId int      `json:"passengerId,omitempty" bson:"PassengerId,omitempty"`
	Survived    bool     `json:"survived,omitempty" bson:"Survived,omitempty"`
	Title       string   `json:"title,omitempty" bson:"Title,omitempty"`
	FirstName   string   `json:"firstName,omitempty" bson:"FirstName,omitempty"`
	MiddleName  string   `json:"middleNameName,omitempty" bson:"MiddleNameName,omitempty"`
	LastName    string   `json:"lastName,omitempty" bson:"LastName,omitempty"`
	Sex         Sex      `json:"sex,omitempty"  bson:"Sex,omitempty"`
	Age         int      `json:"age,omitempty"  bson:"Age,truncate"`
	TripData    TripData `json:"tripData"`
}
