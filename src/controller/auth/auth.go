package auth

import (
	"casamento_api/src/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func GenerateToken(guest model.Guest) (string, error) {
	secret := "secret"

	claims := jwt.MapClaims{
		"id":           guest.ID,
		"name":         guest.Name,
		"phone_number": guest.PhoneNumber,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("error signing token")
	}

	return SignedToken, nil
}

func VerifyToken(tokenString string) (*model.Guest, error) {
	secret := "secret"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errors.New("invalid token")
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return &model.Guest{
		ID:          claims["id"].(uuid.UUID),
		Name:        claims["name"].(string),
		PhoneNumber: claims["phone_number"].(string),
	}, nil
}

func removeBarearPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearear ")
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := "secret"
	tokenValue := removeBarearPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errors.New("invalid token")
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	idString, ok := claims["id"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	guest := model.Guest{
		ID:          id,
		Name:        claims["name"].(string),
		PhoneNumber: claims["phone_number"].(string),
	}

	fmt.Println(guest)
}
