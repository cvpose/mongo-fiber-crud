package database

import (
	"os"
	"testing"

	"github.com/cvpose/crud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	_, cleanup, err := testutil.SetupMongoContainer()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestInitDatabase(t *testing.T) {
	assert.NotPanics(t, func() {
		InitDatabase()
	})
}

// func TestInitDatabaseWithoutAppName(t *testing.T) {
// 	// Remove APP_NAME só para esse teste
// 	os.Unsetenv("APP_NAME")

// 	assert.NotPanics(t, func() {
// 		InitDatabase()
// 	})

// 	// Restaura se quiser
// 	os.Setenv("APP_NAME", "test_app")
// }
