package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRepository_Create(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	model := &TestModel{Name: "Alice", Age: 30}
	err := repo.Create(ctx, model)
	assert.NoError(t, err)
	assert.NotEmpty(t, model.ID.Hex())
}

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

func TestRepository_Update(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	model := &TestModel{Name: "Alice", Age: 30}
	repo.Create(ctx, model)
	update := bson.M{"age": 31}
	updated := &TestModel{}
	err := repo.Update(ctx, model.ID.Hex(), update, updated)
	assert.NoError(t, err)
	assert.Equal(t, 31, updated.Age)
}

func TestRepository_Count(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	count, err := repo.Count(ctx, bson.M{"name": "Alice"})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestRepository_FindWithPagination(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	var paged []TestModel
	total, err := repo.FindWithPagination(ctx, bson.M{"name": "Alice"}, 1, 10, &paged)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, paged, 1)
}

func TestRepository_FindOne(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	one := &TestModel{}
	err := repo.FindOne(ctx, bson.M{"name": "Alice"}, one)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", one.Name)
}

func TestRepository_UpdateMany(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	modCount, err := repo.UpdateMany(ctx, bson.M{"name": "Alice"}, bson.M{"age": 32})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), modCount)
}

func TestRepository_DeleteMany(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	repo.Create(ctx, &TestModel{Name: "Alice", Age: 30})
	delCount, err := repo.DeleteMany(ctx, bson.M{"name": "Alice"})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), delCount)
}

func TestRepository_Delete(t *testing.T) {
	// cleanup := setupTestDB(t)
	// defer cleanup()
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	model := &TestModel{Name: "Alice", Age: 30}
	repo.Create(ctx, model)
	repo.Delete(ctx, model.ID.Hex())
	err := repo.Delete(ctx, model.ID.Hex())
	assert.Error(t, err)
}
