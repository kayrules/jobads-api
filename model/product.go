package model

import (
	"errors"
	"log"

	"../config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	// Product struct
	Product struct {
		ID    bson.ObjectId `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
		Code  string        `json:"code" form:"code" bson:"code" valid:"required"`
		Name  string        `json:"name" form:"name" bson:"name" valid:"required"`
		Price int           `json:"price" form:"price" bson:"price" valid:"int,required"`
	}
)

// ProductIndexing to create indices
// ----------------------------------------------------------------------
func ProductIndexing() {
	c := DB.Copy().DB(config.DbName).C(config.ProductsCollection)
	err := c.EnsureIndex(mgo.Index{
		Key:    []string{"code"},
		Unique: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// CreateProduct Crud
// ----------------------------------------------------------------------
func CreateProduct(product *Product) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.ProductsCollection)

	numRows, err := c.Find(bson.M{"code": product.Code}).Count()
	if err != nil {
		return err
	}
	if numRows > 0 {
		return errors.New("Product Code already exists")
	}

	if err = c.Insert(product); err != nil {
		return errors.New("Creating Product failed")
	}

	return err
}

// ListProduct cRud
// ----------------------------------------------------------------------
func ListProduct() (results []*Product, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.ProductsCollection)

	err = c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}

// SelectProductByID cRud
// ----------------------------------------------------------------------
func SelectProductByID(id bson.ObjectId) (result *Product, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.ProductsCollection)

	err = c.FindId(id).One(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// UpdateProduct crUd
// ----------------------------------------------------------------------
func UpdateProduct(id bson.ObjectId, update *Product) (result *Product, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.ProductsCollection)

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

// DeleteProduct cruD
// ----------------------------------------------------------------------
func DeleteProduct(id bson.ObjectId) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.ProductsCollection)

	err = c.RemoveId(id)
	if err != nil {
		return err
	}

	return err
}
