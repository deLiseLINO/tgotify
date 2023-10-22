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
	UserByID(id uint) (*models.User, error)
	UserByName(name string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id uint) error
	ChangePassword(uid uint, pass string) error
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
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}

	// Retrieve the user information based on the user ID.
	user, err := a.DB.UserByID(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	err := c.BindJSON(&userI)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash the user's password for security.
	hashed, err := hashPass(userI.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Create a new user with the provided information.
	user := models.User{Name: userI.Name, Password: hashed}

	// Add the user to the database.
	err = a.DB.CreateUser(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate a token for the user.
	token, err := authentication.GenerateToken(user.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// DeleteUser handles the deletion of a user's account.
func (a *UserApi) DeleteUser(c *gin.Context) {
	// Extract the user ID from the JWT token in the request context.
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}

	// Call the DeleteUser method from the userDB interface to delete the user by their ID.
	err := a.DB.DeleteUser(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// Signin handles the user sign-in process.
func (a *UserApi) Signin(c *gin.Context) {
	// Parse the JSON input data for sign-in.
	var userI userInput
	err := c.BindJSON(&userI)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Retrieve the user's information by their name.
	user, err := a.DB.UserByName(userI.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Compare the provided password with the stored password hash.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userI.Password))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate a token for the authenticated user.
	token, err := authentication.GenerateToken(user.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (a *UserApi) ChangePassword(c *gin.Context) {
	uid := c.GetUint("user_id")
	if uid == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fetchuid)
		return
	}
	pass := c.PostForm("pass")
	if pass == "" {
		newErrorResponse(c, http.StatusBadRequest, "missing pass parameter")
	}

	// Hash the user's password for security.
	hashed, err := hashPass(pass)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = a.DB.ChangePassword(uid, hashed)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func hashPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
