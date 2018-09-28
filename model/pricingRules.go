package model

import (
	"errors"
	"log"

	"../config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	// PricingRules struct
	PricingRules struct {
		ID            bson.ObjectId `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
		CustomerID    string        `json:"customer_id" form:"customer_id" bson:"customer_id" valid:"required"`
		ProductCode   string        `json:"product_code" form:"product_code" bson:"product_code" valid:"required"`
		Type          string        `json:"type" form:"type" bson:"type" valid:"required,in(deal|discount)"`
		DealBuy       int           `json:"deal_buy" form:"deal_buy" bson:"deal_buy" valid:"-"`
		DealPriceOf   int           `json:"deal_priceof" form:"deal_priceof" bson:"deal_priceof" valid:"-"`
		DiscountBuy   int           `json:"discount_buy" form:"discount_buy" bson:"discount_buy" valid:"-"`
		DiscountPrice int           `json:"discount_price" form:"discount_price" bson:"discount_price" valid:"-"`
	}
)

// PricingRulesIndexing to create indices
// ----------------------------------------------------------------------
func PricingRulesIndexing() {
	c := DB.Copy().DB(config.DbName).C(config.PricingRulesCollection)
	err := c.EnsureIndex(mgo.Index{
		Key:    []string{"customer_id", "product_code", "type"},
		Unique: false,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// CreatePricingRules Crud
// ----------------------------------------------------------------------
func CreatePricingRules(rules *PricingRules) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

	numRows, err := c.Find(bson.M{
		"customer_id":  rules.CustomerID,
		"product_code": rules.ProductCode,
	}).Count()
	if err != nil {
		return err
	}
	if numRows > 0 {
		return errors.New("Rules already exists for the same customer & product_code")
	}

	if err = c.Insert(rules); err != nil {
		return errors.New("Creating Rules failed")
	}

	return err
}

// ListPricingRules cRud
// ----------------------------------------------------------------------
func ListPricingRules() (results []*PricingRules, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

	err = c.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}

// SelectPricingRulesByID cRud
// ----------------------------------------------------------------------
func SelectPricingRulesByID(id bson.ObjectId) (result *PricingRules, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

	err = c.FindId(id).One(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// SelectPricingRulesByCustomerID cRud
// ----------------------------------------------------------------------
func SelectPricingRulesByCustomerID(id string) (results []*PricingRules, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

	err = c.Find(bson.M{"customer_id": id}).All(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}

// UpdatePricingRules crUd
// ----------------------------------------------------------------------
func UpdatePricingRules(id bson.ObjectId, update *PricingRules) (result *PricingRules, err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

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

// DeletePricingRules cruD
// ----------------------------------------------------------------------
func DeletePricingRules(id bson.ObjectId) (err error) {
	db := DB.Clone()
	defer db.Close()
	c := db.DB(config.DbName).C(config.PricingRulesCollection)

	err = c.RemoveId(id)
	if err != nil {
		return err
	}

	return err
}
