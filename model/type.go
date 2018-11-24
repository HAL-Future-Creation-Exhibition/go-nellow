package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Icon                 string        `json:"icon"`
	Name                 string        `json:"name"`
	ProvidingDestination `json:"providing_destination"`
	// NellowID             bson.ObjectId `json:"nellow_id" bson:"nellow_id"`
}

type Users []User

type ProvidingDestination struct {
	PName string `json:"name"`
	PRate uint   `json:"rate"`
}

type Nellow struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
