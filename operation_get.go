package repository

import (
	"context"
	"errors"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetByID retrieves a document by its ID
func (r *Repository) GetByID(ctx context.Context, id string, result mgm.Model) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid object ID format")
	}

	err = r.collection.FindByID(objectID, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("document not found")
		}
		return err
	}

	return nil
}

// GetAll retrieves all documents matching the filter
func (r *Repository) GetAll(ctx context.Context, filter bson.M, results interface{}, opts ...*options.FindOptions) error {
	if filter == nil {
		filter = bson.M{}
	}

	err := r.collection.SimpleFind(results, filter, opts...)
	if err != nil {
		return err
	}

	return nil
}
