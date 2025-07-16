package repository

import "github.com/kamva/mgm/v3"

// Training represents a training entity stored in MongoDB.
//
// It embeds mgm.DefaultModel to provide default fields such as ID, CreatedAt, and UpdatedAt.
// Fields:
//   - Name:        The name of the training.
//   - Description: A description of the training.
//   - ImageURL:    The URL of the image associated with the training.
type Training struct {
	*mgm.DefaultModel `bson:",inline"`
	Name              string `json:"name" bson:"name"`
	Description       string `json:"description" bson:"description"`
	ImageURL          string `json:"image_url" bson:"image_url"`
}

func (model *Training) CollectionName() string {
	return "trainings" // Specify the MongoDB collection name
}

// Export the model symbol and its name
var Model Training
var ModelName = "trainings" // Plural, for route
