package model

import (
	"github.com/globalsign/mgo/bson"
)

type (
	// Purchase struct
	Purchase struct {
		CustomerID bson.ObjectId `json:"customer_id,omitempty" form:"customer_id,omitempty" bson:"customer_id,omitempty"`
		Classic    int           `json:"classic" form:"classic" bson:"classic" valid:"int,optional"`
		Standout   int           `json:"standout" form:"standout" bson:"standout" valid:"int,optional"`
		Premium    int           `json:"premium" form:"premium" bson:"premium" valid:"int,optional"`
		Total      int           `json:"total,omitempty" form:"total,omitempty" bson:"total,omitempty"`
	}
)
