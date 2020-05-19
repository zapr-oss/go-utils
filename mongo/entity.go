package mongo

import (
	"fmt"
	"log"
	"github.com/zapr-oss/go-utils/mongo/config"
	"gopkg.in/mgo.v2"
	"net/url"
)

type Entity struct {
	DB      *mgo.Database
	Session *mgo.Session
}

//InitMongoDB creates a connection to MongoDB using the provided mongo config.
//Returns pointers to the MongoDB object, and the open session to it.
//(Mongo) -> (*MongoDatabase, *Session)
func Connect(mongoConfig mongoconfig.Config) (*Entity, error) {
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

	return &Entity{
		DB:      database,
		Session: session,
	}, nil
}

func (e *Entity) Close() {
	e.Session.Close()
}
