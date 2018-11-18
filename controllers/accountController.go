package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/cards/models"
	"fmt"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AccountContext struct {
	Db *mongo.Database
}


func (c *AccountContext) CreateUser(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*models.User)
	if((body.Email == "") || (body.Password == "")) {
		models.WriteError(w, models.ErrUserMissingData)
		return
	}

	// encrypt password
	// TODO Add Pepper to the password string - http://0xdabbad00.com/2015/04/23/password_authentication_for_go_web_servers/
	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 0)
	if err != nil {
		models.WriteError(w, models.ErrInternalServer)
		return
	}

	body.Password = string(bytes)

	rJson := models.BasicJSONReturn{}
	rJson.ReturnType = "registration"
	w.Header().Set("Content-Type", "application/vnd.api+json")

	// 1. Does user exist?
	repo := models.UserRepo{c.Db.Collection("users")}
	fmt.Println("body:", body)
	if repo.UserExist(body) {
		// a user already exists
		w.WriteHeader(701)
		rJson.ReturnStatus =  models.ErrUserAlreadyExists.Title
		rJson.Payload = models.ErrUserAlreadyExists.Detail
	}else{
		// no user found. Create user record
		id, err := repo.Create(body)
		if err != nil {
			w.WriteHeader(500)
			rJson.ReturnStatus = models.ErrInternalServer.Title
			rJson.Payload = models.ErrInternalServer.Detail
		}else{
			fmt.Println("objectid.ObjectID ID:", id.String())
			w.WriteHeader(201)
			rJson.ReturnStatus = "success"
			rJson.Payload = id.String()
		}

	}

	json.NewEncoder(w).Encode(rJson)
}


func (c *AccountContext) Login(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*models.User)
	if((body.Email == "") || (body.Password == "")) {
		models.WriteError(w, models.ErrUserMissingData)
		return
	}

	fmt.Println("Inside Login 1")
	//params := r.URL.Query()
	//email := params.Get("email")
	//password := params.Get("password")

	repo := models.UserRepo{c.Db.Collection("users")}
	user, err := repo.Login(body.Email, body.Password)
	if err != nil {
		if ((err.Error() == models.ErrUserNotFound.Title) || (err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password")) {
			models.WriteError(w, models.ErrUserNotFound)
			return
		} else {
			models.WriteError(w, models.ErrInternalServer)
			return
		}
	}

	fmt.Println("Inside Login 2")
	// create a JWT with claims
	// Create the Claims
	//noinspection ALL
	claims := models.MyCustomClaims{
		user.Data.UUID,
		jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Second * 3600 * 24).Unix(),
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cardsApi",
			NotBefore: 0,
			Subject:   "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(models.GetSecret()))
	fmt.Printf("%v %v", signedString, err)

	fmt.Println("cardsApi userController LoginHandler: JWT Signed jwtString = ", signedString)

	fmt.Println("Inside Login 3")
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(signedString)

}

func (c *AccountContext) UserProfile(w http.ResponseWriter, r *http.Request) {
	// verify token  and extract claims before doing anything!
	jwtToken := r.Header.Get("Token")
	claims, ok := models.ExtractClaims(jwtToken)
	if !ok {
		models.WriteError(w, models.ErrUserTokenRejected)
		return
	}
	// Get user's Profile
	repo := models.ProfileRepo{c.Db.Collection("profiles")}
	profile, err := repo.GetPublicProfile(claims["id"].(string))
	if err != nil{
		models.WriteError(w, models.ErrUserNotFound)
	}

fmt.Println("Profile:",profile)

}