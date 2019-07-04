package repository

import (
	"github.com/elgiavilla/mc_user/models"
	"github.com/elgiavilla/mc_user/users"
	"github.com/juju/mgosession"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	pool *mgosession.Pool
	db   string
}

func NewMongoRepo(p *mgosession.Pool, db string) users.Repository {
	return &MongoRepository{
		pool: p,
		db:   db,
	}
}

func (m *MongoRepository) Find(id models.ID) (*models.User, error) {
	result := models.User{}
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C("users")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoRepository) FindAll() ([]*models.User, error) {
	var d []*models.User
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C("users")
	err := coll.Find(nil).All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoRepository) Store(b *models.User) (models.ID, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C("users")
	err := coll.Insert(b)
	if err != nil {
		return models.ID(0), err
	}
	return b.ID, nil
}

func (m *MongoRepository) Delete(id models.ID) error {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C("users")
	return coll.Remove(bson.M{"_id": id})
}
