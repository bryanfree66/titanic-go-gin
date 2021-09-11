package model

type Passenger struct {
	Id                      string  `json:"id,omitempty" bson:"_id,omitempty"`
	PassengerId             string  `json:"passengerId,omitempty" bson:"PassengerId,omitempty"`
	Survived                bool    `json:"survived,omitempty" bson:"Survived,omitempty"`
	PassengerClass          int     `json:"passengerClass,omitempty" bson:"PClass,omitempty"`
	Name                    string  `json:"name,omitempty" bson:"Name,omitempty"`
	Sex                     string  `json:"sex,omitempty"  bson:"Sex,omitempty"`
	Age                     int     `json:"age,omitempty"  bson:"Age,truncate"`
	SiblingsOrSpousesAboard int     `json:"siblingsOrSpousesAboard,omitempty" bson:"SibSp,omitempty"`
	ParentsOrChildrenAboard int     `json:"parentsOrChildrenAboard,omitempty" bson:"Parch,omitempty"`
	Ticket                  int     `json:"ticket,omitempty" bson:"Ticket,omitempty"`
	Fare                    float32 `json:"fare,omitempty" bson:"Fare,truncate"`
	Cabin                   string  `json:"cabin,omitempty" bson:"Cabin,omitempty"`
	Embarked                string  `json:"embarked,omitempty" bson:"Embarked,omitempty"`
}
