package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QrHistory struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`   
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`      
	Name      string             `json:"name" bson:"name"`
	URL       string             `json:"url" bson:"url"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
  }
  