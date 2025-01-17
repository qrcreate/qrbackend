package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
    ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`         
    Username     string             `json:"username" bson:"username"`        
    Email        string             `json:"email" bson:"email"`             
    Password     string             `json:"password,omitempty" bson:"password,omitempty"` 
    PasswordHash string             `json:"passwordhash,omitempty" bson:"passwordhash,omitempty"` 
    CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`    
    UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`     
}

type QrHistory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
	Type        string             `bson:"type" json:"type"`
	Content     string             `bson:"content" json:"content"`
	Name        string             `bson:"name" json:"name"`
	QrImage     string             `bson:"qrImage" json:"qrImage"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}