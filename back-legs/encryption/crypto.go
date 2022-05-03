package encryption

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func EncryptText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), 8)
	if err != nil {
		return "err", err
	}
	return string(hashedText), nil
}

func Compare(text string, encryptedText string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedText), []byte(text)); err != nil {
		return false
	}
	return true
}

func CreateJWT(username string) string {
	var jwtKey = []byte("my_secret_key")
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(15000 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return ""
	}
	return tokenString
}

// https://www.sohamkamani.com/golang/jwt-authentication/

func ParseJWT(token string) (bool, error) {
	claims := &Claims{}
	var jwtKey = []byte("my_secret_key")
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}
	if !tkn.Valid {
		// w.WriteHeader(http.StatusUnauthorized)
		return false, errors.New("Invalid token")
	}
	return true, nil
}
