package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

// BaseStruct struct
type BaseStruct struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	CreatedAt string        `json:"createdAt" bson:"createdAt"`
	UpdatedAt string        `json:"updatedAt" bson:"updatedAt"`
}
