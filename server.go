package main

import (
	"API-E-commerce/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type ID_Producto struct {
	Id_Producto int `json:"id_Producto"`
}

type Inicio_Sesion struct {
	Acceso_Valido bool `json:"acceso_valido"`
}

func main() {
	fmt.Print("Bienvenido \n")
	initServer()
}

func initServer() {
	router := gin.Default()
	router.POST("/api/clientes/iniciar_sesion", postLogin)
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
	var user models.Cliente
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Login(user)
		c.IndentedJSON(http.StatusCreated, user)
	}
}

func postCompra(c *gin.Context) {
	var det models.Detalle
	var comp models.Compra
	if err := c.BindJSON(&comp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Addcompra(comp)
		models.Adddetalle(det)
		c.IndentedJSON(http.StatusCreated, comp)
	}
}

func postProduct(c *gin.Context) {
	var prod models.Producto
	var id_prod ID_Producto
	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.Addproduct(prod)
		id_prod.Id_Producto = models.GetIdByName(prod.Nombre)
		c.IndentedJSON(http.StatusCreated, id_prod)
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
	var id_prod ID_Producto
	id := c.Param("id")
	product := models.GetProductbyid(id)
	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {
		models.Delete(id)
		var err error
		id_prod.Id_Producto, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		c.IndentedJSON(http.StatusOK, id_prod)
	}
}

func getStatistics(c *gin.Context) {

}

func putProductByID(c *gin.Context) {
	var prod models.Producto

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		var id_prod ID_Producto
		id_str := c.Param("id")
		id_int, err := strconv.Atoi(id_str)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		id_prod.Id_Producto = id_int
		prod.Id_Producto = id_int
		models.Putproduct(prod)
		c.IndentedJSON(http.StatusCreated, id_prod)
	}
}
