package main

type ConfigFile struct {
	Port		int `yaml:"port,omitempty"`
	SigningKey	string `yaml:"signing_key,omitempty"`
	MongoDbServer	string `yaml:"mongodb_server,omitempty"`
	MongoDbDatabase	string `yaml:"mongodb_database,omitempty"`
}