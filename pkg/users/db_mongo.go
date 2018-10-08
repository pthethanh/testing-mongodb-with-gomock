package users

import (
	"fmt"

	"github.com/golovers/testing-mongodb-with-gomock/pkg/env"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

var _ Repository = &mongoDB{}

type mongoDB struct {
	conf *mongCfg
	conn *mgo.Session
}

type mongCfg struct {
	DbURI      string `envconfig:"USER_DB_URI" default:"localhost:27017"`
	DbName     string `envconfig:"USER_DB_NAME" default:"mongo_test"`
	DbUser     string `envconfig:"USER_DB_USERNAME"`
	DbPassword string `envconfig:"USER_DB_PASSWORD"`
}

// NewMongoDB creates a new article repository backed by a given Mongo server,
// authenticated with given credentials.
func NewMongoDB() (*mongoDB, error) {
	var cfg mongCfg
	env.LoadEnvConfig(&cfg)

	conn, err := mgo.Dial(cfg.DbURI)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}
	if cfg.DbUser != "" {
		cred := &mgo.Credential{
			Username: cfg.DbUser,
			Password: cfg.DbPassword,
		}
		if cred != nil {
			if err := conn.Login(cred); err != nil {
				return nil, err
			}
		}
	}
	return &mongoDB{
		conf: &cfg,
		conn: conn,
	}, nil
}

func (mongo *mongoDB) GetUsers() ([]*User, error) {
	s := mongo.conn.Clone()
	defer s.Close()
	var users []*User
	err := s.DB(mongo.conf.DbName).C("users").Find(bson.D{}).All(&users)
	return users, err
}

func (mongo *mongoDB) AddUser(u User) error {
	s := mongo.conn.Clone()
	defer s.Close()
	return s.DB(mongo.conf.DbName).C("users").Insert(u)
}
