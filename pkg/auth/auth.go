package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/luccasbarros/the-service/cmd/api/config"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
  // should be changed (add id, etc)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(config.Env.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  if err != nil {
    return "", err
  }

  return string(bytes), nil
}

func ComparePassword(hashedPassword, password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  return err != nil
}
