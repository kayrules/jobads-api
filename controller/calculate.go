package controller

import (
	"net/http"

	"../model"
	"github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

// Calculate godocs
// ----------------------------------------------------------------------
// @tags Product
// @Summary Calculate purchase
// @Description Calculate purchase
// @Accept  json
// @Produce  json
// @Param Body body model.Product true " "
// @Success 200 {object} model.Product
// @Failure 400 {object} echo.HTTPError
// @Router /calculate/{customer_id} [post]
// ----------------------------------------------------------------------
func Calculate(c echo.Context) (err error) {
	headers := c.Response().Header()
	headers.Set("Access-Control-Allow-Origin", "*")
	headers.Set("Access-Control-Allow-Headers", "Authorization")
	headers.Set("Access-Control-Allow-Methods", "GET, PATCH, PUT, POST, DELETE, OPTIONS")

	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is an invalid ObjectID")
	}
	id := bson.ObjectIdHex(c.Param("id"))

	purchase := &model.Purchase{
		CustomerID: id,
	}
	if err = c.Bind(purchase); err != nil {
		return err
	}

	_, err = govalidator.ValidateStruct(purchase)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// get customer rules
	var rules []*model.PricingRules
	rules, err = model.SelectPricingRulesByCustomerID(id.Hex())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	eligiblities := map[string]*model.PricingRules{}
	for _, r := range rules {
		eligiblities[r.ProductCode] = r
	}

	// get product lists
	var products []*model.Product
	products, err = model.ListProduct()
	if err != nil {
		return err
	}
	basePrices := map[string]int{}
	for _, p := range products {
		basePrices[p.Code] = p.Price
	}

	total := 0
	// classic ads
	if purchase.Classic > 0 {
		eg := eligiblities["classic"]
		if eg != nil {
			total += eligiblity(eg, purchase.Classic, basePrices["classic"])
		} else {
			// normal pricing
			total += purchase.Classic * basePrices["classic"]
		}
	}

	// standout ads
	if purchase.Standout > 0 {
		eg := eligiblities["standout"]
		if eg != nil {
			total += eligiblity(eg, purchase.Standout, basePrices["standout"])
		} else {
			// normal pricing
			total += purchase.Standout * basePrices["standout"]
		}
	}

	// premium ads
	if purchase.Premium > 0 {
		eg := eligiblities["premium"]
		if eg != nil {
			total += eligiblity(eg, purchase.Premium, basePrices["premium"])
		} else {
			// normal pricing
			total += purchase.Premium * basePrices["premium"]
		}
	}

	purchase.Total = total
	// log.Println(total)
	return c.JSON(http.StatusOK, purchase)
}

func eligiblity(eg *model.PricingRules, buy int, basePrice int) (total int) {
	// if rule type=deal
	if eg.Type == "deal" {
		whole := buy / eg.DealBuy
		residue := buy % eg.DealBuy
		total = (whole * eg.DealPriceOf * basePrice) + (residue * basePrice)
		// if rule type=discount
	} else if eg.Type == "discount" && buy >= eg.DiscountBuy {
		total = buy * eg.DiscountPrice
		// not eligible = normal pricing
	} else {
		total = buy * basePrice
	}
	return total
}
