package database

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDatabase initializes the MongoDB connection and sets up mgm as the ODM.
//
// It reads the MongoDB URI from the MONGO_URI environment variable and the database name
// from the MONGO_DATABASE environment variable. Optionally, it reads the application name
// from the APP_NAME environment variable (defaults to "training-api" if not set).
//
// If any required environment variable is missing or the connection fails, the application will log a fatal error and exit.
func InitDatabase() {
	var uri string
	var databaseName string
	var appName string
	var err error

	if uri = os.Getenv("MONGO_URI"); uri == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	if databaseName = os.Getenv("MONGO_DATABASE"); databaseName == "" {
		log.Fatal("MONGO_DATABASE environment variable not set")
	}

	if appName = os.Getenv("APP_NAME"); appName == "" {
		log.Println("APP_NAME environment variable not set, using default 'training-api'")
		appName = "training-api"
	}

	clientOptions := options.Client().ApplyURI(uri).SetAppName(appName)

	if err = mgm.SetDefaultConfig(nil, databaseName, clientOptions); err != nil {
		log.Fatalf("Failed to set mgm default config: %v", err)
	}

	log.Println("Connected to MongoDB with mgm!")
}
