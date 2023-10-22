package handlers

import (
	"net/http"
	authentication "tgotify/api/handlers/auth"
	models "tgotify/storage"

	"github.com/gin-gonic/gin"
)

// ClientDB is an interface for database operations related to clients.
type ClientDB interface {
	CreateClient(client models.Client) error
	GetUserByID(uint) (*models.User, error)
}

type ClientApi struct {
	DB ClientDB
}

// ClientInput is a struct used to parse JSON input for creating a client.
type ClientInput struct {
	Name  string `json:"name" binding:"required"`
	Token string `json:"token" binding:"required"`
}

// CreateClient is a handler function for creating a new client.
func (a *ClientApi) CreateClient(c *gin.Context) {
	// Extract the user ID from the JWT token in the request context.
	uid, err := authentication.ExtractTokenID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Parse JSON input to create a new client.
	var cl ClientInput
	err = c.BindJSON(&cl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create a new client object.
	client := models.Client{
		Name:   cl.Name,
		Token:  cl.Token,
		UserID: uid,
	}

	// Call the database method to create the client.
	err = a.DB.CreateClient(client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
