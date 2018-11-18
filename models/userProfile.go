package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"fmt"
	"time"
	"context"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)


type ControlContet struct {
	Db mongo.Database
}

type ProfileRepo struct {
	Coll *mongo.Collection
}

type PublicProfileResource struct {
	Data PublicProfile `json:"data"`
}

type (
	Profile struct {
		Id          objectid.ObjectID    `json:"_id" bson:"_id,omitempty"`
		FName       string        `json:"fName" bson:"fName"`
		LName       string        `json:"lName" bson:"lName"`
		DisplayName string        `json:"displayName" bson:"displayName"`
		UUID        string        `json:"uuid" bson:"uuid"`
	}

	PublicProfile struct {
		FName       string        `json:"fName" bson:"fName"`
		LName       string        `json:"lName" bson:"lName"`
		DisplayName string        `json:"displayName" bson:"displayName"`
	}

)


func (r *ProfileRepo) GetPublicProfile(uuid string) (PublicProfileResource, error) {
	result := PublicProfileResource{}
	//result := bson.NewDocument()
	filter := bson.NewDocument(bson.EC.String("hello", "world"))
	err := r.Coll.FindOne(context.Background(),filter).Decode(result)
	if err != nil {
		fmt.Println(time.Now(), " userProfile.go GetPublicProfile 001: Error: ", ErrInternalServer.Title, " ", err)
		return result, err
	}
	return result, nil
}