package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("learn_oo")

func CreateIdToken(id uint, email string, emailVerified bool) (string, error) {
	jwtIdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":            id,
			"email":         email,
			"emailVerified": emailVerified,
			"exp":           time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	idToken, err := jwtIdToken.SignedString(secretKey)
	return idToken, err
}

func CreateIdTokenFromEmail(email string) (string, error) {
	jwtIdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":         email,
			"exp":           time.Now().Add(time.Minute * 5).Unix(),
		},
	)

	idToken, err := jwtIdToken.SignedString(secretKey)
	return idToken, err
}

func VerifyToken(idToken string) (jwt.MapClaims, error) {
	jwtIdToken, err := jwt.Parse(idToken, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtIdToken.Valid {
		return nil, errors.New("INVALID_TOKEN")
	}

	claims, _ := jwtIdToken.Claims.(jwt.MapClaims)

	return claims, nil
}

func CreateRefreshToken(id uint) (string, error) {
	jwtIdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
		},
	)
	idToken, err := jwtIdToken.SignedString(secretKey)
	return idToken, err
}

func VerifyRefreshToken(refreshToken string) (jwt.MapClaims, error) {
	jwtRefreshToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtRefreshToken.Valid {
		return nil, errors.New("INVALID_REFRESH_TOKEN")
	}

	claims, _ := jwtRefreshToken.Claims.(jwt.MapClaims)

	return claims, nil
}

