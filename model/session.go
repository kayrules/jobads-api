package model

import (
	"log"

	"../config"
	"github.com/globalsign/mgo"
)

// DB global session
var DB *mgo.Session

func init() {
	var err error
	DB, err = mgo.Dial(config.DbHost)
	if err != nil {
		log.Fatal(err)
	}

	ProductIndexing()
	CustomerIndexing()
	PricingRulesIndexing()
}
