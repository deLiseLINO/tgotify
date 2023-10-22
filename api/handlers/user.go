package handlers

import (
	"net/http"
	authentication "tgotify/api/handlers/auth"
	models "tgotify/storage"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// userDB interface defines the methods for interacting with the user database.
type userDB interface {
	GetUserByID(id uint) (*models.User, error)
	UserByName(name string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id uint) error
}

// UserApi is a handler for user-related operations.
type UserApi struct {
	DB userDB
}

// userInput represents the expected JSON input structure for creating or signing in a user.
type userInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pass" binding:"required"`
}

func (a *UserApi) User(c *gin.Context) {
	// Extract the user ID from the JWT token in the request context.
	userID, err := authentication.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get user ID from token",
		})
		return
	}

	// Retrieve the user information based on the user ID.
	user, err := a.DB.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to fetch user from the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Name,
	})
}

// CreateUser handles the creation of a new user.
func (a *UserApi) CreateUser(c *gin.Context) {
	// Parse the JSON input data.
	var userI userInput
	err := c.ShouldBindJSON(&userI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing parameters",
		})
		return
	}

	// Hash the user's password for security.
	hash, err := bcrypt.GenerateFromPassword([]byte(userI.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to hash the password",
		})
		return
	}

	// Create a new user with the provided information.
	user := models.User{Name: userI.Name, Password: string(hash)}

	// Add the user to the database.
	err = a.DB.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create the user",
		})
		return
	}

	// Generate a token for the user.
	token, err := authentication.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond with the generated token and a "200 OK" status.
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// DeleteUser handles the deletion of a user's account.
func (a *UserApi) DeleteUser(c *gin.Context) {
	// Extract the user ID from the JWT token in the request context.
	userID, err := authentication.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get user ID",
		})
		return
	}

	// Call the DeleteUser method from the userDB interface to delete the user by their ID.
	err = a.DB.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to delete the user from the database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Signin handles the user sign-in process.
func (a *UserApi) Signin(c *gin.Context) {
	// Parse the JSON input data for sign-in.
	var userI userInput
	err := c.ShouldBindJSON(&userI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing parameters",
		})
		return
	}

	// Retrieve the user's information by their name.
	user, err := a.DB.UserByName(userI.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// Compare the provided password with the stored password hash.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userI.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// Generate a token for the authenticated user.
	token, err := authentication.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
