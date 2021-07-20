package database

import (
	"os"

	"gopkg.in/mgo.v2"
)

type Mongo struct {
	Posts *mgo.Collection
}

func CreateConn() *Mongo {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}

	db := session.DB("Jobbot")

	return &Mongo{
		Posts: db.C("posts"),
	}

}
