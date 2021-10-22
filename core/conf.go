package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"strconv"
	"strings"
)

type APIConfiguration struct {
	Debug           bool   `required:"false" default:"false" envconfig:"debug"`
	ApplicationPort int    `required:"false" default:"8080" envconfig:"appPort"`
	ApplicationHost string `required:"false" default:"localhost"`
	DBPort          int    `required:"true" default:"5432" envconfig:"DatabasePort"`
	DBHost          string `required:"true" default:"localhost"`
	DBName          string `required:"true" default:"golang_api" envconfig:"dbName"`
	DBUser          string `required:"true" default:"postgres"`
	DBPass          string `required:"true" default:"postgres"`
}

func (config *APIConfiguration) GetHTTPAddr() string {
	return fmt.Sprintf("%s:%d", config.ApplicationHost, config.ApplicationPort)
}

func (config *APIConfiguration) GetDatabaseConnectionURI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName,
	)
}

func (config *APIConfiguration) GetApplicationListenURL() string {
	builder := &strings.Builder{}
	builder.WriteString(config.ApplicationHost)
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(config.ApplicationPort))
	return builder.String()
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
