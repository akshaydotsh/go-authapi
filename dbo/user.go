package dbo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/theakshaygupta/go-authapi/models"
)

func (db *Database) InsertUser(user models.User) error {
	err := db.DB.C("users").Insert(user)
	return err
}
func (db *Database) FindUser(query bson.M) (models.User, error) {
	var user models.User
	err := db.DB.C("users").Find(query).One(&user)
	return user, err
}
func (db *Database) GetUsers(query bson.M) ([]models.User, error) {
	var users []models.User
	err := db.DB.C("users").Find(query).All(users)
	return users, err
}
func (db *Database) UpdateUser(query bson.M, update bson.M) (models.User, error) {
	return models.User{}, nil
}
func (db *Database) DeleteUser(query bson.M) error {
	return nil
}
