package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gidsi/SpaceApiSpec/v013"
	"log"
	"time"
)

type spaceApiWithTimestamp struct {
	Data spaceapi_spec.Root
	Timestamp int64
}

func writeSpaceData(data spaceapi_spec.Root) {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spacedata")
	err = c.Insert(
		spaceApiWithTimestamp{
			data,
			time.Now().Unix(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func readSpacedata() spaceapi_spec.Root {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spacedata")
	result := spaceApiWithTimestamp{}
	err = c.Find(bson.M{}).Sort("-timestamp").One(&result)
	if err != nil {
		panic(err)
	}

	return result.Data
}

func checkToken(token string) bool {
	if(token == "") {
		return false
	}

	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("token")
	result := Token{}
	err = c.Find(bson.M{ "token": token }).One(&result)
	if err != nil {
		return false
	}

	return result.Token != ""
}

func readToken() []Token {
	panic(config)
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("token")
	result := []Token{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		panic(err)
	}

	return result
}


func writeToken(token string) {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("token")
	err = c.Insert(Token{ token })
	if err != nil {
		log.Fatal(err)
	}
}