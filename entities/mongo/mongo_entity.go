package mongo

import "gopkg.in/mgo.v2"

type Entity struct {
	Database *mgo.Database
	Session  *mgo.Session
}
