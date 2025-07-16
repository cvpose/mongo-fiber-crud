package repository

import (
	"os"
	"testing"

	"github.com/cvpose/repository/testutil"
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
)

type TestModel struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name"`
	Age              int    `bson:"age"`
}

func TestMain(m *testing.M) {
	_, cleanup, err := testutil.SetupMongoContainer()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestRepository_New(t *testing.T) {
	repo := New(&TestModel{})

	assert := assert.New(t)
	assert.NotNil(repo, "Repository should not be nil")
	assert.NotNil(repo.collection, "Collection should not be nil")
	assert.Equal("test_models", repo.collection.Name(), "Collection name should be 'test_models'")
}

// test using implementation of CollectionName interface
type OtherTestModel struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name"`
	Age              int    `bson:"age"`
}

func (m *OtherTestModel) CollectionName() string {
	return "othertestmodel"
}

func TestRepository_NewOtherTestModel(t *testing.T) {
	repo := New(&OtherTestModel{})

	assert := assert.New(t)
	assert.NotNil(repo, "Repository should not be nil")
	assert.NotNil(repo.collection, "Collection should not be nil")
	assert.Equal("othertestmodel", repo.collection.Name(), "Collection name should be 'othertestmodel'")
}
