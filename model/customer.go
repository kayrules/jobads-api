package model

import (
	"errors"
	"log"

	"../config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	// Customer struct
	Customer struct {
		ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name string        `json:"name" bson:"name" valid:"required"`
	}
)

// CustomerIndexing to create indices
// ----------------------------------------------------------------------
func CustomerIndexing() {
	c := DB.Copy().DB(config.DbName).C(config.CustomersCollection)
	err := c.EnsureIndex(mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// CreateCustomer Crud
// ----------------------------------------------------------------------
func CreateCustomer(customer *Customer) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	numRows, err := c.Find(bson.M{"name": customer.Name}).Count()
	if err != nil {
		return err
	}
	if numRows > 0 {
		return errors.New("Customer already exists")
	}

	if err = c.Insert(customer); err != nil {
		return errors.New("Creating Customer failed")
	}

	return err
}

// ListCustomer cRud
// ----------------------------------------------------------------------
func ListCustomer() (results []*Customer, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	err = c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}

// SelectCustomerByID cRud
// ----------------------------------------------------------------------
func SelectCustomerByID(id bson.ObjectId) (result *Customer, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	err = c.FindId(id).One(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// SearchCustomer cRud
// ----------------------------------------------------------------------
func SearchCustomer(params map[string]bson.M) (results []*Customer, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	err = c.Find(params).All(&results)
	if err != nil {
		return nil, err
	}

	log.Println(results)

	return results, err
}

// UpdateCustomer crUd
// ----------------------------------------------------------------------
func UpdateCustomer(id bson.ObjectId, update *Customer) (result *Customer, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	err = c.UpdateId(id, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	err = c.FindId(id).One(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// DeleteCustomer cruD
// ----------------------------------------------------------------------
func DeleteCustomer(id bson.ObjectId) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.CustomersCollection)

	err = c.RemoveId(id)
	if err != nil {
		return err
	}

	return err
}
