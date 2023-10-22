package handlers

import (
	"net/http"
	authentication "tgotify/api/handlers/auth"
	models "tgotify/storage"

	"github.com/gin-gonic/gin"
)

// MessageSender is an interface for sending messages.
type MessageSender interface {
	SendMessage(token string, chatID uint, message string) error
}

// MessageDb is an interface for database operations related to messages.
type MessageDb interface {
	EnabledClients(uid uint) ([]models.Client, error)
}

type MessageApi struct {
	Sender MessageSender
	DB     MessageDb
}

// CreateMessage is a handler function for creating and sending a message.
func (a *MessageApi) CreateMessage(c *gin.Context) {
	// Extract the 'message' parameter from the HTTP POST request form.
	message := c.PostForm("message")
	if message == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing 'message' parameter",
		})
		return
	}

	// Extract the user ID from the JWT token in the request context.
	uid, err := authentication.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get user ID",
		})
		return
	}

	// Get a list of enabled clients for the user.
	clients, err := a.DB.EnabledClients(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get clients",
		})
		return
	}

	// Iterate over the enabled clients and their associated chats to send the message.
	for _, client := range clients {
		for _, chat := range client.Chats {
			if err := a.Sender.SendMessage(client.Token, chat.ChatID, message); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to send message",
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}
