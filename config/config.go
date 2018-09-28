package config

import (
	"log"
	"os"
)

// var public settings
var (
	Env    string
	DbHost string
	DbName string
	Port   string
)

func init() {

	Env = os.Getenv("ENV")
	DbHost = os.Getenv("DB_HOST")
	DbName = os.Getenv("DB_NAME")
	Port = os.Getenv("PORT")

	if Env == "" {
		log.Fatal("cannot find ENV from Env")
	}
	if DbHost == "" {
		log.Fatal("cannot find DB_HOST from Env")
	}
	if DbName == "" {
		log.Fatal("cannot find DB_NAME from Env")
	}
	if Port == "" {
		log.Fatal("cannot find PORT from Env")
	}
}

//IsProduction to check whether Environment is production
func IsProduction() bool {
	return Env == "production"
}
