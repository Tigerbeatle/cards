package models

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"time"
)

type BasicJSONReturn struct {
	ReturnType   string          `json:"ReturnType"        bson:"ReturnType"`
	ReturnStatus string          `json:"ReturnStatus"        bson:"ReturnStatus"`
	Payload      string          `json:"payLoad"        bson:"payLoad"`
}

type Configuration struct {
	Secret string
	Pepper string
}

/*
func ValidateToken(tokenStr string) (interface{}, error){
	hmacSecretString := GetSecret()
	hmacSecret := []byte(hmacSecretString)

	jwtToken := r.Header.Get("Token")
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSecret()), nil
	})

	return token, err

}


*/

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := GetSecret()
	hmacSecret := []byte(hmacSecretString)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func  GetSecret() string {
	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(time.Now()," Configuration.go GetSecret 001: Error: ",ErrInternalServer.Title, " ", err)
		return ""
	}
	return configuration.Secret
}

func  GetPepper() string {
	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(time.Now()," Configuration.go GetPepper 001: Error: ",ErrInternalServer.Title, " ", err)
		return ""
	}
	return configuration.Pepper
}
