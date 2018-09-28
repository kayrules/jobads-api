package main

import (
	config "./config"

	"./controller"
	_ "./docs"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
)

// @title Job Ads Checkout System
// @version 1.0
// @description This is API spec for Job Ads Checkout System.

// @contact.name Kay-Rules
// @contact.url https://kayrules.com/
// @contact.email kayrules@gmail.com

// @host https://kayrules.com/
// @BasePath /

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())

	// product routes
	e.GET("/products", controller.ProductListing)
	e.POST("/product/create", controller.ProductCreate)
	e.GET("/product/:id", controller.ProductSelectByID)
	e.PUT("/product/:id", controller.ProductUpdate)
	e.DELETE("/product/:id", controller.ProductDelete)

	// customer routes
	e.GET("/customers", controller.CustomerSearch)
	e.POST("/customer/create", controller.CustomerCreate)
	e.GET("/customer/:id", controller.CustomerSelectByID)
	e.PUT("/customer/:id", controller.CustomerUpdate)
	e.DELETE("/customer/:id", controller.CustomerDelete)

	// rules routes
	e.GET("/rules", controller.PricingRulesListing)
	e.POST("/rule/create", controller.PricingRulesCreate)
	e.GET("/rule/:id", controller.PricingRulesSelectByID)
	e.GET("/rule/customer/:id", controller.PricingRulesSelectByCustomerID)
	e.PUT("/rule/:id", controller.PricingRulesUpdate)
	e.DELETE("/rule/:id", controller.PricingRulesDelete)

	// calculation routes
	e.POST("/calculate/:id", controller.Calculate)

	//Routes for specs
	e.GET("/specs/*", echoSwagger.WrapHandler)

	// Start server
	e.Server.Addr = ":" + config.Port
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
