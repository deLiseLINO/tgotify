package handlers

import (
	"fmt"
	"net/http"
	models "tgotify/storage"

	"github.com/gin-gonic/gin"
)

// MessageSender is an interface for sending messages.
type MessageSender interface {
	SendMessage(token string, chatID uint, message string) error
}

// MessageDB is an interface for database operations related to messages.
type MessageDB interface {
	EnabledClientsForUser(uid uint) ([]models.Client, error)
}

type MessageAPI struct {
	Sender MessageSender
	DB     MessageDB
}

// CreateMessage is a handler function for creating and sending a message.
func (a *MessageAPI) CreateMessage(c *gin.Context) {
	type Message struct {
		Text string `json:"text"`
	}
	message := Message{}
	// Extract the 'message' parameter from the HTTP POST request form.
	err := c.BindJSON(&message)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if message.Text == "" {
		newErrorResponse(c, http.StatusBadRequest, "missing `text` parameter")
		return
	}

	// Extract the user ID from the JWT token in the request context.
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}

	// Get a list of enabled clients for the user.
	clients, err := a.DB.EnabledClientsForUser(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var count int
	// Iterate over the enabled clients and their associated chats to send the message.
	for _, client := range clients {
		for _, chat := range client.Chats {
			count++
			if err := a.Sender.SendMessage(client.Token, chat.ChatID, message.Text); err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	c.JSON(http.StatusOK, statusResponse{
		fmt.Sprintf("Succesufully send to %d chats", count),
	})
}
