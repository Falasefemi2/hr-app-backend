package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OnboardingWorkflow struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Steps       string             `bson:"steps" json:"steps"` // JSON string representing steps
}
