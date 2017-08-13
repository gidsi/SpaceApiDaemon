package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"
	"strings"
)

var config = ConfigFile{
	ApiPort:               8080,
	SigningKey:            "",
	MongoConnectionString: "localhost",
	MongoCollection:       "spaceApi",
	AllowedOrigins:        []string { "http://localhost" },
}

type ConfigFile struct {
	ApiPort               int `yaml:"apiPort,omitempty"`
	SigningKey            string `yaml:"signingKey,omitempty"`
	MongoConnectionString string `yaml:"mongoConnectionString,omitempty"`
	MongoCollection       string `yaml:"mongoCollection,omitempty"`
	AllowedOrigins        []string `yaml:"allowedOrigins,omitempty"`
}

func initConfig() {
	data, _ := ioutil.ReadFile("/etc/spaceapidaemon/config.yaml")
	yaml.Unmarshal(data, &config)

	flag.IntVar(&config.ApiPort, "apiPort", config.ApiPort, "Port of the api")
	flag.StringVar(&config.SigningKey, "signingKey", config.SigningKey, "Signing key for access token")
	flag.StringVar(&config.MongoConnectionString, "mongoConnectionString", config.MongoConnectionString, "Mongodb connection string")
	flag.StringVar(&config.MongoCollection, "mongoCollection", config.MongoCollection, "Mongodb collection")
	allowedOriginsFlag := flag.String("allowedOrigins", strings.Join(config.AllowedOrigins, ","), "Comma seperated list of allowed origins")

	flag.Parse()

	config.AllowedOrigins = strings.Split(*allowedOriginsFlag, ",")

	if config.SigningKey == "" {
		panic("Signing key not provided, please add a key with a value for \"signing_key\" to your config.yaml")
	}
}