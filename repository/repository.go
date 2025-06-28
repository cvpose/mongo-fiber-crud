package repository

import (
	"github.com/cvpose/crud/database"
	"github.com/kamva/mgm/v3"
)

// Repository implements CRUD operations for any model that implements mgm.Model
type Repository struct {
	collection *mgm.Collection
}

// New creates a new instance of Repository for the specified model type
func New(model mgm.Model) *Repository {
	database.InitDatabase()

	return &Repository{
		collection: mgm.Coll(model),
	}
}
