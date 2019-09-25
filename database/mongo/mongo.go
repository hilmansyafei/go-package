package mongo

import (
	"errors"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// BaseStruct struct
type BaseStruct struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	CreatedAt string        `json:"createdAt" bson:"createdAt"`
	UpdatedAt string        `json:"updatedAt" bson:"updatedAt"`
}

// Mongo : hold global variabel
type Mongo struct {
	DB         *mgo.Database
	Collection *mgo.Collection
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

// MongoMock : hold global variabel
type MongoMock struct {
	CollectionsReturn    map[string]interface{}
	InterfaceReturn      interface{}
	InterfaceReturnArray []interface{}
	Data                 string
}

// MongoProvider : interface
type MongoProvider interface {
	Create(collection string, update interface{}, mdl interface{}) (bson.ObjectId, *mgo.ChangeInfo, error)
	Update(collection string, query bson.M, update interface{}) error
	GetAll(collection string, mdl *[]interface{}) error
	GetByID(collection string, ID bson.ObjectId, mdl *interface{}) error
	Get(collection string, query bson.M, mdl *[]interface{}) error
	GetOne(collection string, query bson.M, mdl *interface{}) error
	Delete(collection string, query bson.M) error
	DeleteID(collection string, ID bson.ObjectId) error
	DeleteAll(collection string, query bson.M) (*mgo.ChangeInfo, error)
	Find(collection string, query bson.M, mdl *[]interface{}, pagingQuery PagingQuery) error
	UpdateApply(collection string, query bson.M, change mgo.Change, doc *map[string]interface{}) (*mgo.ChangeInfo, error)
	Pipe(collection string, query []bson.M, mdl *[]interface{}) error
}

// New is create mysql client
func New(cfg Configuration) (*Mongo, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{cfg.Host},
		Timeout:  60 * time.Second,
		Database: cfg.Database,
		Username: cfg.User,
		Password: cfg.Password,
	}
	session, err := mgo.DialWithInfo(info)
	// Check for error
	if err != nil {
		return &Mongo{}, err
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

// UpdateApply : update data with mgp.change
func (m *Mongo) UpdateApply(collection string, query bson.M, change mgo.Change, doc *map[string]interface{}) (*mgo.ChangeInfo, error) {
	Col := m.DB.C(collection)
	return Col.Find(query).Apply(change, doc)
}

// GetAll record with primary key, return unique result
func (m *Mongo) GetAll(collection string, mdl *[]interface{}) error {
	Col := m.DB.C(collection)
	return Col.Find(nil).All(mdl)
}

// GetByID : Get record with condition, return unique result
func (m *Mongo) GetByID(collection string, ID bson.ObjectId, mdl *interface{}) error {
	Col := m.DB.C(collection)
	return Col.FindId(ID).One(mdl)
}

// Get : Get list record
func (m *Mongo) Get(collection string, query bson.M, mdl *[]interface{}) error {
	Col := m.DB.C(collection)
	return Col.Find(query).All(mdl)
}

// GetOne : Get one record
func (m *Mongo) GetOne(collection string, query bson.M, mdl *interface{}) error {
	Col := m.DB.C(collection)
	return Col.Find(query).One(mdl)
}

// Find : Get records with paging query
func (m *Mongo) Find(collection string, query bson.M, mdl *[]interface{}, pagingQuery PagingQuery) error {
	Col := m.DB.C(collection)
	return Col.Find(query).Sort(pagingQuery.Sort).Skip(pagingQuery.Offset).Limit(pagingQuery.Limit).All(mdl)
}

// Pipe : Function to join 1 or more collection.
func (m *Mongo) Pipe(collection string, query []bson.M, mdl *[]interface{}) error {
	Col := m.DB.C(collection)
	pipe := Col.Pipe(query)
	return pipe.All(mdl)
}

// ------------------- MOCK ------------------ //

// Create : Insert new record
func (m *MongoMock) Create(collection string, update interface{}, mdl interface{}) (bson.ObjectId, *mgo.ChangeInfo, error) {
	return "", nil, nil
}

// Delete : delete collection
func (m *MongoMock) Delete(collection string, query bson.M) error {
	return nil
}

// DeleteID : delete collection
func (m *MongoMock) DeleteID(collection string, ID bson.ObjectId) error {
	return nil
}

// DeleteAll : delete collection
func (m *MongoMock) DeleteAll(collection string, query bson.M) (*mgo.ChangeInfo, error) {
	return nil, nil
}

// Update all fields
func (m *MongoMock) Update(collection string, query bson.M, update interface{}) error {
	return nil
}

// UpdateApply : update data with mgp.change
func (m *MongoMock) UpdateApply(collection string, query bson.M, change mgo.Change, doc *map[string]interface{}) (*mgo.ChangeInfo, error) {
	return nil, nil
}

// GetAll record with primary key, return unique result
func (m *MongoMock) GetAll(collection string, mdl *[]interface{}) error {
	*mdl = m.InterfaceReturnArray
	if m.InterfaceReturnArray == nil {
		return errors.New("error")
	}
	return nil
}

// GetByID : mock function
func (m *MongoMock) GetByID(collection string, ID bson.ObjectId, mdl *interface{}) error {
	*mdl = m.InterfaceReturn
	if m.InterfaceReturn == nil {
		return errors.New("error")
	}
	return nil
}

// Get : Get list record
func (m *MongoMock) Get(collection string, query bson.M, mdl *[]interface{}) error {
	*mdl = m.InterfaceReturnArray
	if m.InterfaceReturnArray == nil {
		return errors.New("error")
	}
	return nil
}

// GetOne : Get one record
func (m *MongoMock) GetOne(collection string, query bson.M, mdl *interface{}) error {
	if len(m.CollectionsReturn) > 0 {
		for table, response := range m.CollectionsReturn {
			if collection == table {
				if response == nil {
					return errors.New("error")
				} else if response == "not found" {
					return errors.New("not found")
				}
				*mdl = response
			}
		}
	}
	return nil
}

// Find : Get records with paging query
func (m *MongoMock) Find(collection string, query bson.M, mdl *[]interface{}, pagingQuery PagingQuery) error {
	*mdl = m.InterfaceReturnArray
	if m.InterfaceReturnArray == nil {
		return errors.New("error")
	}
	return nil
}

// Pipe : Function to join 1 or more collection.
func (m *MongoMock) Pipe(collection string, query []bson.M, mdl *[]interface{}) error {
	return nil
}
