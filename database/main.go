package database

import (
	"gopkg.in/mgo.v2"
	"os"
)

type Mongo struct {
	Posts *mgo.Collection
}

func CreateConn() *Mongo{
	session, err := mgo.Dial(os.Getenv(key:"MONGO_URL"))
	if err != nil {
		panic(err)
	}

	db := session.DB(name: "Jobbot")

	return &Mongo{
		Posts: db.C(name : "posts"),
	}

}
