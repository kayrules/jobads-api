package controller

import (
	"net/http"
	"strings"

	"../model"
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

// CustomerCreate godocs
// ----------------------------------------------------------------------
// @tags Customer
// @Summary Create customer
// @Description create new customer
// @Accept  json
// @Produce  json
// @Param Body body model.Customer true " "
// @Success 200 {object} model.Customer
// @Failure 400 {object} echo.HTTPError
// @Router /customer/create [post]
// ----------------------------------------------------------------------
func CustomerCreate(c echo.Context) (err error) {
	customer := &model.Customer{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(customer); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = model.CreateCustomer(customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

// CustomerSearch godocs
// ----------------------------------------------------------------------
// @tags Customer
// @Summary Search Customer
// @Description Show specific customer based on search
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Customer
// @Failure 400 {object} echo.HTTPError
// @Router /customers [get]
// ----------------------------------------------------------------------
func CustomerSearch(c echo.Context) (err error) {
	var filter = map[string]bson.M{}
	qName := strings.ToLower(c.QueryParam("name"))
	if qName != "" {
		filter["name"] = bson.M{"$regex": bson.RegEx{Pattern: qName, Options: "i"}}
	}

	var results []*model.Customer
	results, err = model.SearchCustomer(filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// CustomerSelectByID godocs
// ----------------------------------------------------------------------
// @tags Customer
// @Summary Select Customer by ID
// @Description Show specific customer based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Customer
// @Failure 400 {object} echo.HTTPError
// @Router /customer/{id} [get]
// ----------------------------------------------------------------------
func CustomerSelectByID(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	var result *model.Customer
	result, err = model.SelectCustomerByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// CustomerUpdate godocs
// ----------------------------------------------------------------------
// @tags Customer
// @Summary Update Customer by ID
// @Description Update specific customer based on selected ID
// @Accept  json
// @Produce  json
// @Param Body body model.Customer true " "
// @Success 200 {object} model.Customer
// @Failure 400 {object} echo.HTTPError
// @Router /customer/{id} [put]
// ----------------------------------------------------------------------
func CustomerUpdate(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	customer := new(model.Customer)
	if err = c.Bind(customer); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result *model.Customer
	result, err = model.UpdateCustomer(id, customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// CustomerDelete godocs
// ----------------------------------------------------------------------
// @tags Customer
// @Summary Delete Customer by ID
// @Description Remove specific customer based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Customer
// @Failure 400 {object} echo.HTTPError
// @Router /customer/{id} [delete]
// ----------------------------------------------------------------------
func CustomerDelete(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	err = model.DeleteCustomer(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := map[string]string{
		"status":  "success",
		"message": "Selected customer has been deleted",
	}

	return c.JSON(http.StatusOK, msg)
}
