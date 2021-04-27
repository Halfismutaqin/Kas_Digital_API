package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hadi.com/controllers"
	"hadi.com/models"
)

func main() {

	r := gin.Default()

	//MODEL
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "API Manajemen KAS Digital"})
	})

	r.GET("/pemasukan", controllers.PemasukanTampil)
	r.POST("/pemasukan", controllers.PemasukanTambah)
	r.PUT("/pemasukan/:id_income", controllers.PemasukanUbah)
	r.DELETE("/pemasukan/:id_income", controllers.PemasukanHapus)

	r.GET("/pengeluaran", controllers.PengeluaranTampil)
	r.POST("/pengeluaran", controllers.PengeluaranTambah)
	r.PUT("/pengeluaran/:id_spending", controllers.PengeluaranUbah)
	r.DELETE("/pengeluaran/:id_spending", controllers.PengeluaranHapus)

	r.Run()
}
