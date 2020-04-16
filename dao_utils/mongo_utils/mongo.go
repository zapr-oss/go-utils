package mongo_utils

import (
	"bitbucket.org/zapr/go-utils/entities/mongo"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"net/url"
)

/*
InitMongoDB creates a connection to MongoDB using the provided mongo config.

Returns pointers to the MongoDB object, and the open session to it.

(Mongo) -> (*MongoDatabase, *Session)
*/
func Connect(mongoConfig mongo.Config) (*mongo.Entity, error) {
	mongoUrl := fmt.Sprintf("%s:%d", mongoConfig.Host, mongoConfig.Port)
	if mongoConfig.UserName != "" {
		mongoUrl = fmt.Sprintf("mongodb://%s:%s@%s/%s", mongoConfig.UserName,
			url.QueryEscape(mongoConfig.Password), mongoUrl, mongoConfig.Database)
	}

	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Println("Error while connecting to MongoDB.", err)
		return nil, err
	}

	database := session.DB(mongoConfig.Database)
	return &mongo.Entity{
		Database: database,
		Session:  session,
	}, nil
}
