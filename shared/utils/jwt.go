package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("learn_oo")

func CreateIdToken(id string, authorID *string, emailVerified bool) (string, error) {
	jwtIdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":             id,
			"email_verified": emailVerified,
			"author_id":      authorID,
			"exp":            time.Now().Add(time.Hour * 24 * 2).Unix(),
		},
	)

	idToken, err := jwtIdToken.SignedString(secretKey)
	return idToken, err
}

func CreateIdTokenFromEmail(email string) (string, error) {
	jwtIdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Minute * 5).Unix(),
		},
	)

	idToken, err := jwtIdToken.SignedString(secretKey)
	return idToken, err
}

type Claims struct {
	ID            string
	EmailVerified bool
	AuthorID      *string
}

func VerifyIDToken(idToken string) (claims *Claims, isValid bool, err error) {
	jwtIdToken, err := jwt.Parse(idToken, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, false, err
	}

	userClaims := &Claims{}
	jwtClaims, _ := jwtIdToken.Claims.(jwt.MapClaims)

	if id, ok := jwtClaims["id"].(string); !ok {
		return nil, false, errors.New("casting id failed")
	} else {
		userClaims.ID = id
	}

	if authorID, ok := jwtClaims["author_id"].(string); ok {
		userClaims.AuthorID = &authorID
	} else {
		userClaims.AuthorID = nil
	}

	if emailVerified, ok := jwtClaims["email_verified"].(bool); !ok {
		return nil, false, errors.New("casting email verified id failed")
	} else {
		userClaims.EmailVerified = emailVerified
	}

	if !isValid {
		return userClaims, false, nil
	}

	return userClaims, true, nil
}

func VerifyIDTokenFromEmail(idToken string) (email *string, isValid bool, err error) {
	jwtIdToken, err := jwt.Parse(idToken, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	fmt.Println(jwtIdToken.Valid)

	if err != nil {
		return nil, false, err
	}

	if !jwtIdToken.Valid {
		return nil, false, nil
	}

	jwtClaims, _ := jwtIdToken.Claims.(jwt.MapClaims)

	if emailToVerify, ok := jwtClaims["email"].(string); !ok {
		return nil, false, errors.New("casting email failed")
	} else {
		return &emailToVerify, true, nil
	}
}
