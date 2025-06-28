package database

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var mongoC testcontainers.Container

func TestMain(m *testing.M) {
	ctx := context.Background()

	// 🧪 Setup do container do MongoDB
	containerReq := testcontainers.ContainerRequest{
		Image:        "mongo:7",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	var err error
	mongoC, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		fmt.Printf("❌ Erro ao iniciar container: %v\n", err)
		os.Exit(1)
	}

	// 🧪 Configura variáveis de ambiente
	host, _ := mongoC.Host(ctx)
	port, _ := mongoC.MappedPort(ctx, "27017")

	uri := fmt.Sprintf("mongodb://%s:%s", host, port.Port())

	os.Setenv("MONGO_URI", uri)
	os.Setenv("MONGO_DATABASE", "test_db")
	os.Setenv("APP_NAME", "test_app")

	// Roda os testes
	code := m.Run()

	// Finaliza
	_ = mongoC.Terminate(ctx)
	os.Exit(code)
}

func TestInitDatabase(t *testing.T) {
	assert.NotPanics(t, func() {
		InitDatabase()
	})
}

func TestInitDatabaseWithoutAppName(t *testing.T) {
	// Remove APP_NAME só para esse teste
	os.Unsetenv("APP_NAME")

	assert.NotPanics(t, func() {
		InitDatabase()
	})

	// Restaura se quiser
	os.Setenv("APP_NAME", "test_app")
}
