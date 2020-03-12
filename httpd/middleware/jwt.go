package middleware

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func secret() string {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Main func --> Error loading .env file")
	}

	return os.Getenv("SECRET_KEY")
}

// CreateJwt -- create jwt
func CreateJwt(userID uint, email string) (string, error) {
	keys := []byte(secret())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims

	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 730).Local()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(keys)
	return tokenString, err
}

// ValidJwt -- validasi
func ValidJwt(MyToken string) (jwt.MapClaims, error) {
	keys := secret()
	token, err := jwt.Parse(MyToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(keys), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err

}
