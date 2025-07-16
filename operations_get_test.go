package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRepository_GetByID(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	model := &TestModel{Name: "Alice", Age: 30}
	repo.Create(ctx, model)

	fetched := &TestModel{}
	err := repo.GetByID(ctx, model.ID.Hex(), fetched)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", fetched.Name)
}

func TestRepository_GetAll(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	var all []TestModel
	err := repo.GetAll(ctx, bson.M{"name": "Alice"}, &all)
	assert.NoError(t, err)
	assert.Len(t, all, 1)
}

func TestRepository_GetByID_InvalidID(t *testing.T) {
	ctx := context.Background()
	repo := New(&TestModel{})
	fetched := &TestModel{}
	err := repo.GetByID(ctx, "invalid_id", fetched)
	assert.Error(t, err)
	assert.Equal(t, "invalid object ID format", err.Error())
}

func TestRepository_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)
	fetched := &TestModel{}
	// Use a valid but non-existent ObjectID
	fakeID := "507f1f77bcf86cd799439011"
	err := repo.GetByID(ctx, fakeID, fetched)
	assert.Error(t, err)
	assert.Equal(t, "document not found", err.Error())
}

// To test GetAll error, we need to simulate an error from SimpleFind. We'll use a custom collection stub for this.
// type errorCollection struct {
// 	mgm.Collection
// }

// func (e *errorCollection) SimpleFind(results interface{}, filter interface{}, opts ...interface{}) error {
// 	return assert.AnError
// }

// func TestRepository_GetAll_Error(t *testing.T) {
// 	repo := New(&TestModel{})
// 	repo.collection = &errorCollection{} // Inject erroring collection
// 	var all []TestModel
// 	err := repo.GetAll(context.Background(), bson.M{"name": "Alice"}, &all)
// 	assert.Error(t, err)
// }
