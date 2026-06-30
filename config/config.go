package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"

	"redbank/infrastructure/postgres"
	//@ahum: imports
)

// NewConfig loads configuration from environment variables or .env file.
func CheckConfigs() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("No config file found!")
		} else {
			panic(fmt.Errorf("Error happened during loading config file: %e", err))
		}
	}

	secretKey := viper.Get("app.secret_key")
	if _, ok := secretKey.(string); !ok {
		log.Fatal("app.secret_key not set. Please provide it.")
	}

}

func ConfigInfras() error {

	// @ahum:edges.group

	// @ahum:connect.load

	// @ahum:end.connect.load

	// @ahum:end.edges.group

	// @ahum:infras.group

	// @ahum:postgres.load
	err := postgres.Configure()
	if err != nil {
		return err
	}
	// @ahum:end.postgres.load

	// @ahum:redis.load

	// @ahum:end.redis.load

	// @ahum:end.infras.group

	//@ahum: loads

	return nil
}
