package testutil

import (
	"context"
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupMongoContainer() (testcontainers.Container, func(), error) {
	ctx := context.Background()
	containerReq := testcontainers.ContainerRequest{
		Image:        "mongo:7",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		return nil, nil, err
	}
	host, _ := mongoC.Host(ctx)
	port, _ := mongoC.MappedPort(ctx, "27017")
	uri := fmt.Sprintf("mongodb://%s:%s", host, port.Port())

	os.Setenv("MONGO_URI", uri)
	os.Setenv("MONGO_DATABASE", "test_db")
	os.Setenv("APP_NAME", "test_app")

	cleanup := func() {
		_ = mongoC.Terminate(ctx)
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("MONGO_DATABASE")
		os.Unsetenv("APP_NAME")
	}
	return mongoC, cleanup, nil
}
