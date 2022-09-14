package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/clientes/iniciar_sesion", postLogin)
	router.POST("/api/compras", postCompra)
	router.POST("/api/productos", postProduct)
	router.GET("/api/productos", getProduct)
	router.GET("/api/productos/:id", getProductByID)
	router.DELETE("api/productos/:id", delProductByID)
	router.GET("/api/estadisticas", getStatistics)

	router.Run("localhost:5000")
}
