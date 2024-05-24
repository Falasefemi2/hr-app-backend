package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	CompanyID primitive.ObjectID `bson:"company_id" json:"company_id"`
}
