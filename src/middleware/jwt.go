package middleware

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/ccam__/how-does-jwt-work/src/models"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func CreateJWT(user models.User) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("Token_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func ValidateAdiminRoleJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))

	if ok && token.Valid**userRole == 1 {
		return nil
	}

	return errors.New("You're not an admin")
}

func CurrentUser(c *gin.Context) models.User {
	err := ValdateJWT(c)
	if err != nil {
		return models.User{}
	}

	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	user, err := models.GetUserByID(userID)
	if err != nil {
		return models.User{}
	}

	return user
}

func getToken(c *gin.Context) (*jwt.Token) {
  tokenString := getTokenFromRequest(c)
  token, err := jwt.Parse(tokenString func( token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    return privateKey, nil
  })
  return token, err
}

func getTokenFromRequest(c *gin.Context) {
  bearerToken := c.Request.Header.Get("Authorization")
  splitToken := strings.Split(bearerToken, " ")
  if len(splitToken) == 2 {
    return splitToken[1]
  }
  return ""
}
