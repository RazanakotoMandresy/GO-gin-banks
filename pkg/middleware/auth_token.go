package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenManage(token *jwt.Token, c *gin.Context) (string, error) {
	// Generate a JWT token
	// token generate Token
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return "", errors.New("failed to create token")
	}
	// ovaia header bearer token le cookie teo
	// Reponse Anle cookie anle usersi
	return tokenString, nil
}
func RequireAuth(c *gin.Context) {
	// maka anle authorization
	// innutitlisable satir tss ininiona aoa anaty cookie
	// extractio  du string dans extract token
	tokenString, err := ExtarctToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "TOKEN Deja expirer")
			return
		}
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Erreur token invalide")
		return
	}
}
func ExtractTokenUUID(ctx *gin.Context) (string, error) {
	tokenString, err := ExtarctToken(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return "", err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprint(claims["uuid"]), nil
	}
	return "", errors.New("une erreur lors du get de l'uuid")
}
func ExtarctToken(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	bearerTokenQuery := ctx.Query("token")
	if bearerTokenQuery != "" {
		return bearerTokenQuery, nil
	}
	return "", errors.New("pas de bearer token , ni de token query")
}
