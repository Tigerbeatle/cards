package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"


	"github.com/mongodb/mongo-go-driver/bson/objectid"


	//"time"
	//"errors"
	"log"
	"context"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type (

	User struct {
		ID         objectid.ObjectID   `json:"id" bson:"_id,omitempty"`
		Email    string `valid:"email" json:"email"       bson:"email" `
		Password string `json:"password"  bson:"password"`
		UUID     string `json:"uuid"      bson:"uuid"`
	}

)

type UserRepo struct {
	Coll *mongo.Collection
}

type UserResource struct {
	Data User `json:"data"`
}


type UsersCollection struct {
	Data []User `json:"data"`
}


func (r *UserRepo) UserExist(user *User) (bool) {
	// 1. Look to see if user record exists
	// 2. If exists, return true
	// 3. if does not exist, return false
	cursor, err := r.Coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("email", user.Email)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	itemRead := User{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&itemRead)
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	// no user found, return empty string
	return false

}

func (r *UserRepo) Create(user *User) (objectid.ObjectID, error) {

	// 1. get a unique UUID value
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	user.UUID = uuid.String()
	// 2. NOW INSERT THE USER RECORD
	res, err := r.Coll.InsertOne(context.Background(), bson.NewDocument(
		bson.EC.String("email", user.Email),
		bson.EC.String("uuid", user.UUID),
		bson.EC.String("password", user.Password),
	))
	return res.InsertedID.(objectid.ObjectID), err
}


func (r *UserRepo) Login(email string, password string) (UserResource, error) {
	result := UserResource{}

	fmt.Println("	Inside Login A")
	fmt.Println("Email:", email)
	cursor, err := r.Coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("passord", password)),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("	Inside Login B")
	defer cursor.Close(context.Background())
	//result := User{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Data)
		fmt.Println("	Inside Login C")
		fmt.Println("stored Password:", result.Data.Password)
		fmt.Println("stored eamil   :", email)
		fmt.Println("passed Password:", password)
		err = bcrypt.CompareHashAndPassword([]byte(result.Data.Password), []byte(password))
		if err != nil {
			fmt.Println("Found the error")
			return result, err // return err if password doesn't match hashed password stored in db
		}
	}
	fmt.Println("result:", result.Data)
	return result, nil
}