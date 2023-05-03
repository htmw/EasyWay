package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/ksharma67/EasyWay/server/app"
	"github.com/ksharma67/EasyWay/server/config"
)

func main() {
	// Initialize the logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Get the configuration for the server
	config := config.GetConfig()

	// Create a new instance of the app with the given configuration
	app := &app.App{}
	app.Initialize(config)

	// Migrate the database schema to ensure it is up-to-date
	app.DBMigrate()

	// Serve the "dist" directory as static files
	fs := http.FileServer(http.Dir("dist"))

	// Add CORS middleware to handle Cross-Origin Resource Sharing
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	})

	// Set up API endpoints
	http.Handle("/", corsMiddleware.Handler(fs))

	// Print a message indicating the server is running
	fmt.Println("Working server on port:3000")

	// Start the server on the specified port
	app.Run(":3000")
}
