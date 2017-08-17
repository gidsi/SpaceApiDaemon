package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gidsi/SpaceApiSpec/v013"
	"log"
	"time"
)

type spaceAPIWithTimestamp struct {
	Data spaceapi_spec.Root
	Timestamp int64
}

var session *mgo.Session

func getSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(config.MongoConnectionString)
		session.SetMode(mgo.Primary, true)
		if err != nil {
			panic(err)
		}
	}

	return session.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(config.MongoCollection).C(collection)
	return s(c)
}

func writeSpaceData(data spaceapi_spec.Root) {
	session := getSession()

	c := session.DB(config.MongoCollection).C("spaceData")
	err := c.Insert(
		spaceAPIWithTimestamp{
			data,
			time.Now().Unix(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func readLastSpaceData() (spaceapi_spec.Root, error) {
	result := spaceAPIWithTimestamp{}
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{}).Sort("-timestamp").One(&result)
	}

	err := withCollection("spaceData", query)
	if err != nil {
		log.Print("cant find spaceapi data in collection")
	}

	return result.Data, err
}

func readSpaceData() ([]spaceAPIWithTimestamp, error) {
	result := []spaceAPIWithTimestamp{}
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{}).Sort("-timestamp").All(&result)
	}

	err := withCollection("spaceData", query)
	if err != nil {
		log.Print("cant find spaceapi data in collection")
	}

	return result, err
}

func checkToken(token string) bool {
	if token == "" {
		return false
	}

	result := Token{}
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{ "token": token }).One(&result)
	}
	err := withCollection("token", query)
	if err != nil {
		return false
	}

	return result.Token != ""
}

func readToken() []Token {
	result := []Token{}
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{}).All(&result)
	}
	err := withCollection("token", query)
	if err != nil {
		panic(err)
	}

	return result
}


func writeToken(token string) {
	query := func(c *mgo.Collection) error {
		return c.Insert(Token{ token })
	}
	err := withCollection("token", query)
	if err != nil {
		panic(err)
	}
}

func removeToken(token string) {
	query := func(c *mgo.Collection) error {
		return c.Remove(Token{ token })
	}
	err := withCollection("token", query)
	if err != nil {
		panic(err)
	}
}

func writeImportedData(data n39Item) {
	bar := spaceapi_spec.Root{}
	bar.State = &spaceapi_spec.State{
		Lastchange: int(data.Value.Lastchange),
		Open: data.Value.Open,
	}

	query := func(c *mgo.Collection) error {
		return c.Insert(
			spaceAPIWithTimestamp{
				bar,
				data.Value.Lastchange,
			},
		)
	}
	err := withCollection("spaceData", query)
	if err != nil {
		panic(err)
	}
}