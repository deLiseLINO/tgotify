package handlers

import (
	"net/http"
	"strconv"
	models "tgotify/storage"

	"github.com/gin-gonic/gin"
)

type TokensUpdater interface {
	UpdateClients() error
}

// ClientDB is an interface for database operations related to clients.
type ClientDB interface {
	CreateClient(client models.Client) error
	UserByID(uid uint) (*models.User, error)
	Clients(uid uint) (*[]models.ClientResponse, error)
	DeleteClient(uid uint, id uint) error
	UpdateClient(uid uint, client models.ClientResponse) error
}

type ClientAPI struct {
	DB            ClientDB
	TokensUpdater TokensUpdater
}

// ClientInput is a struct used to parse JSON input for creating a client.
type ClientInput struct {
	Name  string `json:"name" binding:"required"`
	Token string `json:"token" binding:"required"`
}

// CreateClient is a handler function for creating a new client.
func (a *ClientAPI) CreateClient(c *gin.Context) {
	// Extract the user ID from the JWT token in the request context.
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}

	// Parse JSON input to create a new client.
	var cl ClientInput
	err := c.BindJSON(&cl)
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

	a.TokensUpdater.UpdateClients()

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (a *ClientAPI) Clients(c *gin.Context) {
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}
	clients, err := a.DB.Clients(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, clients)
}

func (a *ClientAPI) DeleteClient(c *gin.Context) {
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}
	clientIDstr := c.PostForm("id")
	clientID, err := strconv.ParseUint(clientIDstr, 10, 0)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = a.DB.DeleteClient(uid, uint(clientID))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	a.TokensUpdater.UpdateClients()

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (a *ClientAPI) UpdateClient(c *gin.Context) {
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}

	var cl models.ClientResponse
	err := c.BindJSON(&cl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if cl.ID == 0 {
		newErrorResponse(c, http.StatusBadRequest, "missing id parameter")
		return
	}

	if cl.Name == "" && cl.Token == "" && cl.Enabled == "" {
		newErrorResponse(c, http.StatusBadRequest, "at least one parameter is required (name, token, enabled)")
	}

	err = a.DB.UpdateClient(uid, cl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	a.TokensUpdater.UpdateClients()

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
