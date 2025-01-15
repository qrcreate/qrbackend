package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
    ID           primitive.ObjectID `bson:"_id,omitempty"`
    Username     string             `bson:"username"`
    Email        string             `bson:"email"`
	Password     string             `bson:"password"`
    PasswordHash string             `bson:"passwordhash"`  // Hashed password
    CreatedAt    time.Time          `bson:"createdAt"`
    UpdatedAt    time.Time          `bson:"updatedAt"`
}
type QrHistory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
	Type        string             `bson:"type" json:"type"`
	Content     string             `bson:"content" json:"content"`
	Name        string             `bson:"name" json:"name"`
	QrImage     string             `bson:"qrImage" json:"qrImage"`
	IsPermanent bool               `bson:"isPermanent" json:"isPermanent"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}