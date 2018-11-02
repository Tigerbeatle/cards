package models

import (


	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/kelseyhightower/envconfig"
	"log"
	"context"
)

// DB abstraction
type DB struct {
	*mongo.Database
}

func NewMongoDB() *DB {
	// Pull in the configuration.
	var mongoConfig connstring.ConnString
	if err := envconfig.Process("mgo", &mongoConfig); err != nil {
		log.Fatalf("Startup Err")
		//return err
	}

	// set up conn to mongoHelper
	client, err := mongo.NewClientFromConnString(mongoConfig)
	if err != nil { log.Fatal(err) }
	// Now connect to the db
	err = client.Connect(context.TODO())
	if err != nil { log.Fatal(err) }

	db := client.Database(mongoConfig.Database)

	return &DB{db}
}