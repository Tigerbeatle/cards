// Package mongoHelper provides mongoHelper connectivity support.

package mongoHelper

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"log"
	"github.com/mongodb/mongo-go-driver/mongo"
)


func Startup() error {

	// Pull in the configuration.
	var mongoConfig connstring.ConnString
	if err := envconfig.Process("mgo", &mongoConfig); err != nil {
		log.Fatalf("Startup Err")
		return err
	}

	// set up conn to mongoHelper
	client, err := mongo.NewClientFromConnString(mongoConfig)

	if err != nil { log.Fatal(err) }
	err = client.Connect(context.TODO())
	if err != nil { log.Fatal(err) }
return err
}