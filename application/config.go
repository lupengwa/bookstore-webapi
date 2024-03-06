package application

import (
	"bookstore-webapi/internal/platform/db"
	"github.com/spf13/viper"
	"log"
)

type BookStoreApiServiceProperty struct {
	PgDbConfig db.PostgreSQLDbConfig `validate:"required"`
	ServerPort int
}

func LoadConfig() BookStoreApiServiceProperty {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	cfg := BookStoreApiServiceProperty{
		PgDbConfig: db.PostgreSQLDbConfig{
			Url:      viper.GetString("PG_DB_URL"),
			Port:     viper.GetString("PG_DB_PORT"),
			DbName:   viper.GetString("PG_DB_NAME"),
			User:     viper.GetString("PG_DB_USER"),
			Password: viper.GetString("PG_DB_PASSWORD"),
		},
		ServerPort: viper.GetInt("PORT"),
	}

	return cfg
}
