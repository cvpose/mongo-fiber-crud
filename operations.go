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

// Update updates a document by its ID with the provided update fields
func (r *Repository) Update(ctx context.Context, id string, update bson.M, result mgm.Model) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid object ID format")
	}

	// First, get the current document to check if it exists
	err = r.collection.FindByID(objectID, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("document not found")
		}
		return err
	}

	// Update the document
	_, err = r.collection.UpdateByID(ctx, objectID, bson.M{"$set": update})
	if err != nil {
		return err
	}

	// Return the updated document
	err = r.collection.FindByID(objectID, result)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a document by its ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid object ID format")
	}

	// Delete the document directly
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("document not found")
	}

	return nil
}

// Count returns the number of documents matching the filter
func (r *Repository) Count(ctx context.Context, filter bson.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindWithPagination retrieves documents with pagination support
func (r *Repository) FindWithPagination(ctx context.Context, filter bson.M, page, limit int64, results interface{}) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}

	// Calculate skip value
	skip := (page - 1) * limit

	// Set up find options
	findOptions := options.Find().SetSkip(skip).SetLimit(limit)

	// Get documents
	err := r.collection.SimpleFind(results, filter, findOptions)
	if err != nil {
		return 0, err
	}

	// Get total count
	total, err := r.Count(ctx, filter)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// FindOne retrieves a single document matching the filter
func (r *Repository) FindOne(ctx context.Context, filter bson.M, result mgm.Model) error {
	if filter == nil {
		filter = bson.M{}
	}

	err := r.collection.First(filter, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("document not found")
		}
		return err
	}

	return nil
}

// UpdateMany updates multiple documents matching the filter
func (r *Repository) UpdateMany(ctx context.Context, filter bson.M, update bson.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}

	result, err := r.collection.UpdateMany(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

// DeleteMany removes multiple documents matching the filter
func (r *Repository) DeleteMany(ctx context.Context, filter bson.M) (int64, error) {
	if filter == nil {
		return 0, errors.New("filter cannot be empty for delete many operation")
	}

	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
