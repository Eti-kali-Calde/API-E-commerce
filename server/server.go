package server

import (
	"API-E-commerce/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	fmt.Print("Bienvenido \n")
	initServer()
}

func initServer() {
	router := gin.Default()
	//router.POST("/api/clientes/iniciar_sesion", postLogin)
	router.POST("/api/compras", postCompra)
	router.POST("/api/productos", postProduct)
	router.GET("/api/productos", getProduct)
	router.GET("/api/productos/:id", getProductByID)
	router.PUT("/api/productos/:id", putProductByID)
	router.DELETE("api/productos/:id", delProductByID)
	//router.GET("/api/estadisticas", getStatistics)
	router.Run("localhost:5000")
	fmt.Print("iniciando el server")
}

func postLogin(c *gin.Context) {
	fmt.Print("Loegado\n")
}

func postCompra(c *gin.Context) {
	var det models.Detalle
	var comp models.Compra
	if err := c.BindJSON(&comp, &det); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Addcompra(comp)
		models.Adddetalle(det)
		c.IndentedJSON(http.StatusCreated, comp)
	}
}

func postProduct(c *gin.Context) {
	var prod models.Producto
	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Addproduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}

func getProduct(c *gin.Context) {
	product := models.GetProduct()
	if product == nil || len(product) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	product := models.GetProductbyid(id)
	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func delProductByID(c *gin.Context) {
	id := c.Param("id")
	product := models.GetProductbyid(id)
	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {
		models.Delete(id)
		c.IndentedJSON(http.StatusOK, product)
	}
}

func getStatistics(c *gin.Context) {

}

func putProductByID(c *gin.Context) {
	var prod models.Producto
	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Putproduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}
