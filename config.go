package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AwsAccessKeyId       string `env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey   string `env:"AWS_SECRET_ACCESS_KEY"`
	AwsDefaultRegion     string `env:"AWS_DEFAULT_REGION"`
	AwsStorageBucketName string `env:"AWS_STOARGE_BUCKET_NAME"`
}

// Read properties from config.env file
// Command line enviroment variable will overwrite config.env properties
func NewConfig(configFile string) *Config {
	config := Config{}
	err := cleanenv.ReadConfig(configFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	cleanenv.ReadEnv(&config)
	return &config
}
