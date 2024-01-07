package api

import (
	"tgotify/api/handlers"
	authentication "tgotify/api/handlers/auth"
	telegram "tgotify/client"
	"tgotify/storage/postgres"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CreateRouter initializes and configures the Gin router.
func CreateRouter(tg *telegram.Client, db *postgres.Gormdb) {
	// Initialize the Gin router with default settings.
	gin := gin.Default()

	// Create instances of API handlers, passing the Telegram client and the database connection.
	messageHandler := handlers.MessageAPI{Sender: tg, DB: db}
	userHandler := handlers.UserAPI{DB: db}
	clientHandler := handlers.ClientAPI{DB: db, ClientsUpdater: tg}

	// Create a group for authenticated routes and add the JWT authentication middleware.
	auth := gin.Group("")
	auth.Use(authentication.JwtAuthMiddleware())

	// Define authenticated routes with corresponding HTTP methods and handler functions.
	auth.POST("/message", messageHandler.CreateMessage) // Create a new message

	auth.POST("/client", clientHandler.CreateClient)   // Create a new client
	auth.GET("/client", clientHandler.Clients)         // Retrieve all clients info
	auth.DELETE("/client", clientHandler.DeleteClient) // Deletes client
	auth.PUT("/client", clientHandler.UpdateClient)    // Updates clent

	auth.GET("/user", userHandler.User)                     // Retrieve user information
	auth.DELETE("/user", userHandler.DeleteUser)            // Delete a user's account
	auth.POST("/user/password", userHandler.ChangePassword) // Change user's password

	// Define unauthenticated routes.
	gin.POST("/user", userHandler.CreateUser)    // Create a new user
	gin.POST("/user/signin", userHandler.Signin) // User sign-in

	// Retrieve the router port from the configuration using Viper.
	port := viper.GetString("router.port")
	if port == "" {
		logrus.Fatal("unable to fetch port from config")
	}
	// Start the Gin router and listen on the specified port.
	if err := gin.Run(":" + port); err != nil {
		logrus.Fatal(err)
	}
}
