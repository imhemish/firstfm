package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/imhemish/firstradio/repository"
)

func GenerateToken(user repository.User) (string, error) {
	exptime := time.Hour * 200

	claims := jwt.MapClaims{
		"uid": user.UserID,
		"exp": time.Now().Add(exptime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))

}
