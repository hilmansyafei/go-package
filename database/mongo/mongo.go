package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BaseStruct struct
type BaseStruct struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	CreatedAt string        `json:"createdAt" bson:"createdAt"`
	UpdatedAt string        `json:"updatedAt" bson:"updatedAt"`
}

// Mongo : hold global variabel
type Mongo struct {
	DB        *mgo.Database
	Collectio *mgo.Collection
}

// Configuration : hold data config database
type Configuration struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
	Env      string
}

// PagingQuery : hold data paging
type PagingQuery struct {
	Offset int
	Limit  int
	Sort   string
}

// MongoProvider : interface
type MongoProvider interface {
	Create(collection string, update interface{}, mdl interface{}) (bson.ObjectId, *mgo.ChangeInfo, error)
	Update(collection string, query bson.M, update interface{}) error
	GetAll(collection string, mdl *[]interface{}) error
	GetByID(collection string, ID bson.ObjectId, mdl *interface{}) error
	GetOne(collection string, query bson.M, mdl interface{}) error
	Delete(collection string, query bson.M) error
	DeleteID(collection string, ID bson.ObjectId) error
	DeleteAll(collection string, query bson.M) (*mgo.ChangeInfo, error)
	Find(collection string, query bson.M, mdl *[]interface{}, pagingQuery PagingQuery) error
}

// New is create mysql client
func New(cfg Configuration) (*Mongo, error) {
	session, err := mgo.Dial(cfg.Host + "/" + cfg.Database)
	// Check for error
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := session.DB(cfg.Database)
	return &Mongo{DB: db}, err
}

// Create : Insert new record
func (m *Mongo) Create(collection string, update interface{}, mdl interface{}) (bson.ObjectId, *mgo.ChangeInfo, error) {
	Col := m.DB.C(collection)
	id := bson.NewObjectId()
	info, err := Col.UpsertId(id, update)
	return id, info, err
}

// Delete : delete collection
func (m *Mongo) Delete(collection string, query bson.M) error {
	Col := m.DB.C(collection)
	return Col.Remove(query)
}

// DeleteID : delete collection
func (m *Mongo) DeleteID(collection string, ID bson.ObjectId) error {
	Col := m.DB.C(collection)
	return Col.RemoveId(ID)
}

// DeleteAll : delete collection
func (m *Mongo) DeleteAll(collection string, query bson.M) (*mgo.ChangeInfo, error) {
	Col := m.DB.C(collection)
	return Col.RemoveAll(query)
}

// Update all fields
func (m *Mongo) Update(collection string, query bson.M, update interface{}) error {
	Col := m.DB.C(collection)
	return Col.Update(query, update)
}

// GetAll record with primary key, return unique result
func (m *Mongo) GetAll(collection string, mdl *[]interface{}) error {
	Col := m.DB.C(collection)
	return Col.Find(nil).All(mdl)
}

// GetByID : Get record with condition, return unique result
func (m *Mongo) GetByID(collection string, ID bson.ObjectId, mdl *interface{}) error {
	Col := m.DB.C(collection)
	return Col.FindId(ID).One(&mdl)
}

// GetOne : Get one record
func (m *Mongo) GetOne(collection string, query bson.M, mdl interface{}) error {
	Col := m.DB.C(collection)
	Col.Find(query).One(&mdl)
	return nil
}

// Find : Get records with paging query
func (m *Mongo) Find(collection string, query bson.M, mdl *[]interface{}, pagingQuery PagingQuery) error {
	Col := m.DB.C(collection)
	return Col.Find(nil).Sort(pagingQuery.Sort).Skip(pagingQuery.Offset).Limit(pagingQuery.Limit).All(mdl)
}
