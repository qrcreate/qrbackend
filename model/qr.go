package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	GoogleID   string             `bson:"googleId" json:"googleId,omitempty"`
	ProfilePic string             `bson:"profilePic" json:"profilePic"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type QRHistory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
	Type        string             `bson:"type" json:"type"`
	Content     string             `bson:"content" json:"content"`
	Name        string             `bson:"name" json:"name"`
	IsPermanent bool               `bson:"isPermanent" json:"isPermanent"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}