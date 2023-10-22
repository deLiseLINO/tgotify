package authentication

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

// GenerateToken generates a JWT token for a given user ID.
func GenerateToken(userid uint) (string, error) {
	// Retrieve the token lifespan and secret key from the configuration.
	tokenLifespan := viper.GetInt("router.token_lifespan")
	if tokenLifespan == 0 {
		return "", errors.New("unable to fetch token_lifespan from config")
	}
	secretKey := viper.GetString("router.secret_key")
	if secretKey == "" {
		return "", errors.New("unable to fetch secret_key from config")
	}

	// Create JWT claims with user authorization and expiration time.
	claims := jwt.MapClaims{}
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()

	// Create a new JWT token with the claims and sign it with the secret key.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// TokenValid checks the validity of a JWT token in the request context.
func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := parseToken(tokenString)
	if err != nil {
		return err
	}
	id, err := extractTokenID(token)
	c.Set("user_id", id)
	if err != nil {
		return err
	}
	return nil
}

// ExtractToken retrieves a JWT token from the request context.
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID extracts the user ID from a JWT token.
func extractTokenID(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

// parseToken parses a JWT token with the secret key from the configuration.
func parseToken(tokenString string) (*jwt.Token, error) {
	secretKey := viper.GetString("router.secret_key")
	if secretKey == "" {
		return nil, errors.New("unable to fetch secret_key from config")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return token, err
}
