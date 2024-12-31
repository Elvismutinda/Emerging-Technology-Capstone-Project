package auth

import (
	"backend/config"
	"backend/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CustomClaims struct {
	UserData models.User `json:"user-data"`
	jwt.RegisteredClaims
}

func GenerateAuthToken(userData models.User) (string, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return "", err
	}

	secret := cfg.JWTSecretKey
	customClaims := jwt.MapClaims{
		"user-data": userData,
	}

	now := time.Now()
	expiresAt := 48 * time.Hour

	// initialize standard claims
	claims := jwt.MapClaims{
		"jti": uuid.NewString(),          // Id
		"sub": userData.Id,               // Subject
		"iss": "",                        // issuer
		"exp": now.Add(expiresAt).Unix(), // ExpiresAt
		"iat": now.Unix(),                // IssuedAt
		"nbf": now.Unix(),                // NotBefore
	}

	// add any custom claims
	for key, value := range customClaims {
		claims[key] = value
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	// check signing method
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Check if the token is valid and the claims are of the expected type
	if token.Valid {
		return token, nil
	}

	return nil, errors.New("invalid claims or token")
}

// checks if entered password is correct
func ComparePasswords(hashedPwd []byte, plainPwd []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return false, err
	}
	return true, nil
}
