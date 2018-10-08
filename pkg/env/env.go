package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//LoadEnvConfig load configuration from environtment variables
func LoadEnvConfig(t interface{}) {
	if err := envconfig.Process("", t); err != nil {
		log.Fatalf("config: Unable to load config for %T: %s", t, err)
	}
}
