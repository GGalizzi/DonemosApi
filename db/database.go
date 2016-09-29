package db

import (
	"github.com/flioh/DonemosApi/modelo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	sesión          *mgo.Session
	nombreColección string
}

func NewDatabase(sesión *mgo.Session, n string) *Database {
	return &Database{sesión, n}
}

func (db *Database) GetMongoDB() *mgo.Database {
	return db.sesión.DB("donemos")
}

func (db *Database) Colección() *mgo.Collection {
	return db.sesión.DB("donemos").C(db.nombreColección)
}

func (db *Database) Todos(limit int) *mgo.Query {
	return db.Colección().Find(nil).Limit(limit)
}

func (db *Database) Create(m modelo.IModelo) error {
	id := bson.NewObjectId()
	m.SetId(id)
	err := db.Colección().Insert(m)

	return err
}

func (db *Database) Read(hexId string) (m modelo.IModelo, err error) {
	objectId := bson.ObjectIdHex(hexId)
	err = db.Colección().FindId(objectId).One(&m)

	return
}

func (db *Database) Update(hexId string, m modelo.IModelo) error {
	m.SetIdHex(hexId)
	return db.Colección().UpdateId(m.GetId(), m)
}

func (db *Database) Delete(hexId string) error {
	return db.Colección().RemoveId(bson.ObjectIdHex(hexId))
}
