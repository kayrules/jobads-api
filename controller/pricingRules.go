package controller

import (
	"net/http"

	"../model"
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

// PricingRulesCreate godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary Create rule
// @Description create new rule
// @Accept  json
// @Produce  json
// @Param Body body model.PricingRules true " "
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rule/create [post]
// ----------------------------------------------------------------------
func PricingRulesCreate(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.FormValue("customer_id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "customer_id is an invalid ObjectID")
	}
	customerID := c.FormValue("customer_id")

	rule := &model.PricingRules{
		ID:         bson.NewObjectId(),
		CustomerID: customerID,
	}
	if err = c.Bind(rule); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(rule)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = model.CreatePricingRules(rule); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, rule)
}

// PricingRulesListing godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary PricingRules listings
// @Description List all rules
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rules [get]
// ----------------------------------------------------------------------
func PricingRulesListing(c echo.Context) (err error) {
	var results []*model.PricingRules
	results, err = model.ListPricingRules()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// PricingRulesSelectByID godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary Select PricingRules by ID
// @Description Show specific rule based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rule/{id} [get]
// ----------------------------------------------------------------------
func PricingRulesSelectByID(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	var result *model.PricingRules
	result, err = model.SelectPricingRulesByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PricingRulesSelectByCustomerID godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary Select PricingRules by Customer ID
// @Description Show created rules based on selected Customer ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rule/customer/{customer_id} [get]
// ----------------------------------------------------------------------
func PricingRulesSelectByCustomerID(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := c.Param("id")

	var results []*model.PricingRules
	results, err = model.SelectPricingRulesByCustomerID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// PricingRulesUpdate godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary Update PricingRules by ID
// @Description Update specific rule based on selected ID
// @Accept  json
// @Produce  json
// @Param Body body model.PricingRules true " "
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rule/{id} [put]
// ----------------------------------------------------------------------
func PricingRulesUpdate(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	rule := new(model.PricingRules)
	if err = c.Bind(rule); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(rule)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result *model.PricingRules
	result, err = model.UpdatePricingRules(id, rule)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// PricingRulesDelete godocs
// ----------------------------------------------------------------------
// @tags PricingRules
// @Summary Delete PricingRules by ID
// @Description Remove specific rule based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.PricingRules
// @Failure 400 {object} echo.HTTPError
// @Router /rule/{id} [delete]
// ----------------------------------------------------------------------
func PricingRulesDelete(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	err = model.DeletePricingRules(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := map[string]string{
		"status":  "success",
		"message": "Selected rule has been deleted",
	}

	return c.JSON(http.StatusOK, msg)
}
