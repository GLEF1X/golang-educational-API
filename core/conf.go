package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type APIConfiguration struct {
	Debug           bool   `required:"false" default:"false" envconfig:"debug"`
	ApplicationPort int    `required:"false" default:"8080" envconfig:"appPort"`
	DBPort          int    `required:"true" default:"5432" envconfig:"DatabasePort"`
	DBHost          string `required:"true" default:"localhost"`
	DBName          string `required:"true" default:"golang_api" envconfig:"dbName"`
	DBUser          string `required:"true" default:"postgres"`
	DBPass          string `required:"true" default:"postgres"`
}

func (config *APIConfiguration) GetHTTPAddr() string {
	return fmt.Sprintf(":%d", config.ApplicationPort)
}

func (config *APIConfiguration) GetDsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort,
	)
}

func NewAPIConfiguration() *APIConfiguration {
	var apiConfiguration APIConfiguration
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := envconfig.Process("", &apiConfiguration); err != nil {
		log.Fatal(err.Error())
	}
	return &apiConfiguration
}
