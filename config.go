package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var config = ConfigFile{
	Port: 8080,
	SigningKey: "AllYourBase",
	MongoDbServer: "localhost",
	MongoDbDatabase: "spaceApi",
}

type ConfigFile struct {
	Port		int `yaml:"port,omitempty"`
	SigningKey	string `yaml:"signing_key,omitempty"`
	MongoDbServer	string `yaml:"mongodb_server,omitempty"`
	MongoDbDatabase	string `yaml:"mongodb_database,omitempty"`
}

func initConfig() {
	data, _ := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(data, &config)
}