package controller

import (
	"net/http"

	"../model"
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

// ProductCreate godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Create product
// @Description create new product
// @Accept  json
// @Produce  json
// @Param Body body model.Product true " "
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /product/create [post]
// ----------------------------------------------------------------------
func ProductCreate(c echo.Context) (err error) {
	product := &model.Product{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(product); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = model.CreateProduct(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

// ProductListing godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Products listings
// @Description List all products
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /products [get]
// ----------------------------------------------------------------------
func ProductListing(c echo.Context) (err error) {
	var results []*model.Product
	results, err = model.ListProduct()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

// ProductSelectByID godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Select Product by ID
// @Description Show specific product based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /product/{id} [get]
// ----------------------------------------------------------------------
func ProductSelectByID(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	var result *model.Product
	result, err = model.SelectProductByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// ProductUpdate godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Update Product by ID
// @Description Update specific product based on selected ID
// @Accept  json
// @Produce  json
// @Param Body body model.Product true " "
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /product/{id} [put]
// ----------------------------------------------------------------------
func ProductUpdate(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	product := new(model.Product)
	if err = c.Bind(product); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result *model.Product
	result, err = model.UpdateProduct(id, product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// ProductDelete godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Delete Product by ID
// @Description Remove specific product based on selected ID
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /product/{id} [delete]
// ----------------------------------------------------------------------
func ProductDelete(c echo.Context) (err error) {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	err = model.DeleteProduct(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	msg := map[string]string{
		"status":  "success",
		"message": "Selected product has been deleted",
	}

	return c.JSON(http.StatusOK, msg)
}
