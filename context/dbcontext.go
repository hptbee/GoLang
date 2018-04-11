package context

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	hosts      = "ds153948.mlab.com:53948"
	database   = "todos_db"
	username   = "trungtp"
	password   = "trungtp123"
	collection = "todos"
)

func Connect() *mgo.Database {
	infos := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(infos)
	if err != nil {
		log.Fatal(err)
	}

	return session.DB(database)
}

// Find list of todos
