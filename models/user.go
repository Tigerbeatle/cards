package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"


	"github.com/mongodb/mongo-go-driver/bson/objectid"
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

type UsersCollection struct {
	Data []User `json:"data"`
}


/*
func (r *UserRepo) All() (UsersCollection, error) {
	result := UsersCollection{[]User{}}
	coll := db.Collection("users")




	err := r.Coll.Find(nil).All(&result.Data)
	if err != nil {
		fmt.Println(time.Now()," User.go All 001: Error: ",ErrInternalServer.Title, " ", err)
		return result, err
	}

	return result, nil
}
*/