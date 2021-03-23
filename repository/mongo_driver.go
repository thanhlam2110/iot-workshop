package repository

import (
	"github.com/thanhlam/iot-workshop/model"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

// ProfileRepositoryMongo - ProfileRepositoryMongo
type ProfileRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

//NewProfileRepositoryMongo - NewProfileRepositoryMongo
func NewProfileRepositoryMongo(db *mgo.Database, collection string) *ProfileRepositoryMongo {
	return &ProfileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

//<------------------USER------------------>
//SaveUser
func (r *ProfileRepositoryMongo) SaveUser(userSSO *model.UserSSO) error {
	err := r.db.C(r.collection).Insert(userSSO)
	return err
}

//FindByUser
func (r *ProfileRepositoryMongo) FindByUser(username string) (*model.UserSSO, error) {
	var userSSO model.UserSSO
	err := r.db.C(r.collection).Find(bson.M{"username": username}).One(&userSSO)
	if err != nil {
		return nil, err
	}
	return &userSSO, nil
}

//FindThings
func (r *ProfileRepositoryMongo) FindMapThingChanel(thingid, chanelid string) (*model.MapThingChanel, error) {
	var mapThingChanel model.MapThingChanel
	err := r.db.C(r.collection).Find(bson.M{"chanelid": chanelid, "thingid": thingid}).One(&mapThingChanel)
	if err != nil {
		return nil, err
	}
	return &mapThingChanel, nil
}
