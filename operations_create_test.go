package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	ctx := context.Background()
	repo := New(&TestModel{})
	repo.collection.Drop(ctx)

	model := &TestModel{Name: "Alice", Age: 30}
	err := repo.Create(ctx, model)
	assert.NoError(t, err)
	assert.NotEmpty(t, model.ID.Hex())
}
