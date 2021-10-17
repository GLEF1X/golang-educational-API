package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type APIConfiguration struct {
	Debug bool `required:"false" default:"false" envconfig:"debug"`
	Port  int  `required:"false" default:"80" envconfig:"appPort"`
}

func (config *APIConfiguration) GetHTTPAddr() string {
	return fmt.Sprintf(":%d", config.Port)
}

func GetAPIConfiguration() *APIConfiguration {
	var apiConfiguration APIConfiguration
	if err := envconfig.Process("", &apiConfiguration); err != nil {
		log.Fatal(err.Error())
	}
	return &apiConfiguration
}
