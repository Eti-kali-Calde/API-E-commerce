package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Server() {
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
	router.DELETE("api/productos/:id", delProductByID)
	router.GET("/api/estadisticas", getStatistics)
	router.Run("localhost:5000")
	fmt.Print("iniciando el server")
}

func postLogin(c *gin.Context) {
	fmt.Print("Loegado\n")
}

func postCompra(c *gin.Context) {}

func postProduct(c *gin.Context) {}

func getProduct(c *gin.Context) {}

func getProductByID(c *gin.Context) {}

func delProductByID(c *gin.Context) {}

func getStatistics(c *gin.Context) {}
