package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	e := GetEngine()
	e.PUT("/products", ProductUpdateGin)
}

//ProductCreate is
func ProductCreate(w http.ResponseWriter, r *http.Request) {
	var product Product

	if !getRequestBody(w, r, &product) {
		responseError(w, "cannot map request body")
		return
	}

	if !insert(&product) {
		responseError(w, "Cannot insert to database")
		return
	}

	responseSuccess(w, product)
}

func deductPrice(c interface{}) bool {
	return true
}

//ProductUpdate is
func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	var product Product

	if !getRequestBody(w, r, &product) {
		responseError(w, "cannot map request body")
		return
	}

	if !update(&product, func(v interface{}) bool {
		v.(*Product).Price = v.(*Product).Price - 10
		if v.(*Product).Price < 0 {
			return false
		}
		return true
	}) {
		responseError(w, "cannot update due to not enough balance")
		return
	}

	responseSuccess(w, product)
}

//ProductUpdateGin is
func ProductUpdateGin(c *gin.Context) {
	var product Product
	var reqProduct ReqProduct
	c.Bind(&reqProduct)
	product.Code = reqProduct.Code
	if !update(&product, func(v interface{}) bool {
		v.(*Product).Price = v.(*Product).Price - reqProduct.Nominal
		if v.(*Product).Price < 0 {
			return false
		}
		return true
	}) {
		c.JSON(412, ErrorResponse{ACK: false, Message: "not enough balance"})
		return
	}
	c.JSON(200, product)
}
