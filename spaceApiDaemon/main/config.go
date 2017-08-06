package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var config = ConfigFile{
	Port: 8080,
	SigningKey: "",
	MongoDbServer: "database",
	MongoDbDatabase: "spaceApi",
	AllowedOrigins: []string { "http://localhost" },
}

type ConfigFile struct {
	Port		int `yaml:"port,omitempty"`
	SigningKey	string `yaml:"signing_key,omitempty"`
	MongoDbServer	string `yaml:"mongodb_server,omitempty"`
	MongoDbDatabase	string `yaml:"mongodb_database,omitempty"`
	AllowedOrigins	[]string `yaml:"allowed_origins,omitempty"`
}

func initConfig() {
	data, err := ioutil.ReadFile("config.yaml")
	yaml.Unmarshal(data, &config)

	if err != nil {
		panic("could not read config.yaml, will not start.")
	}

	if config.SigningKey == "" {
		panic("Signing key not provided, please add a key with a value for \"signing_key\" to your config.yaml")
	}
}