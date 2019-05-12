package dbo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/theakshaygupta/go-authapi/models"
)

type DatabaseOps interface {
	Close()

	// users
	InsertUser(user models.User) error
	FindUser(query bson.M) (models.User, error)
	GetUsers(query bson.M) ([]models.User, error)
	UpdateUser(query bson.M, update bson.M) (models.User, error)
	DeleteUser(query bson.M) error
}

type Database struct {
	DB      *mgo.Database
	Session *mgo.Session // only for closing connection
}

func Connect() (DatabaseOps, error) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Database{DB: session.DB("go-authapi"), Session: session}, nil
}

func (db *Database) Close() {
	db.Session.Close()
}
