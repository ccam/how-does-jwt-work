package middleware

import (
	"os"

	"gitlab.com/ccam__/how-does-jwt-work/src/models"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func CreateJWT(user models.User) {
}
