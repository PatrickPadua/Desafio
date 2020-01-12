package suggestion

import (
	mgo "gopkg.in/mgo.v2"
	"log"
)

type MongoDB struct {
	Server   string
	Database string
}

var db *mgo.Database

func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (s *Suggestion) Create(COLLECTION string) error {
	err := db.C(COLLECTION).Insert(*s)
	return err
}