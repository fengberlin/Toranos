package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/loehnertz/Toranos/common"
	"github.com/micro/go-log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const RetrieveTokenSecret = "SELECT value FROM secrets WHERE key = $1"
const PasswordFromDb = "SELECT password FROM users WHERE email = $1"

func issueUserToken(email string, password string) (successful bool, token string) {
	var passwordFromDb string
	row := database.QueryRow(PasswordFromDb, email)
	selectError := row.Scan(&passwordFromDb)
	if selectError != nil {
		log.Log(selectError)
		return
	}

	if comparePasswords(passwordFromDb, password) {
		successful, token = createNewToken(
			email,
			common.GetConfigStringByPath(conf, "service-settings", "user-management", "audienceKeyCustomer"),
		)
	}

	return
}

func authenticateUser(token string) (successful bool, email string, role string) {
	successful, email, role = verifyToken(token)

	return
}

func createNewToken(subject string, audience string) (bool, string) {
	// Create a new token
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"aud": audience,
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString(tokenSigningKey)
	if err != nil {
		log.Log(err)
		return false, ""
	}
	return true, tokenString
}

func verifyToken(tokenString string) (successful bool, subject string, audience string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the correct algorithm was used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSigningKey, nil
	})

	// Validate the token itself
	if err != nil {
		log.Log(err)
		successful = false
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Log(err)
		successful = false
	} else if ok && token.Valid {
		successful = true
		subject = claims["sub"].(string)
		audience = claims["aud"].(string)
	}

	return
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Log(err)
		return false
	} else {
		return true
	}
}

func retrieveTokenSecret() ([]byte, error) {
	var tokenSecret string
	row := database.QueryRow(
		RetrieveTokenSecret,
		common.GetConfigStringByPath(conf, "service-settings", "user-management", "tokenSecretTableKey"),
	)
	selectError := row.Scan(&tokenSecret)
	if selectError != nil {
		log.Log(selectError)
		return []byte(""), selectError
	} else {
		return []byte(tokenSecret), nil
	}
}
