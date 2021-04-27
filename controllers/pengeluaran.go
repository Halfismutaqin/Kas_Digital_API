package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hadi.com/models"
)

type PengeluaranInput struct {
	Id_spending string `json:"id_spending"`
	Tanggal     string `json:"tanggal"`
	Jumlah      int    `json:"jumlah"`
	Sumber      string `json:"sumber"`
}

//Tampil data pengeluaran
func PengeluaranTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var keluar []models.Spending
	db.Find(&keluar)
	c.JSON(http.StatusOK, gin.H{"data": keluar})
}

// Tambah data Pengeluaran
func PengeluaranTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/Keluarkan
	var dataInput PengeluaranInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	keluar := models.Spending{
		Id_spending: dataInput.Id_spending,
		Tanggal:     dataInput.Tanggal,
		Jumlah:      dataInput.Jumlah,
		Sumber:      dataInput.Sumber,
	}

	db.Create(&keluar)

	c.JSON(http.StatusOK, gin.H{"data": keluar})
}

// Ubah data Pengeluaran
func PengeluaranUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var Keluar models.Spending
	if err := db.Where("id_spending = ?", c.Param("id_spending")).First(&Keluar).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Pengeluaran tidak ditemukan"})
		return
	}

	//validasi input/Keluarkan
	var dataInput PengeluaranInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&Keluar).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": Keluar})
}

// Hapus data Pengeluaran
func PengeluaranHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var Keluar models.Spending
	if err := db.Where("id_spending = ?", c.Param("id_spending")).First(&Keluar).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Pengeluaran tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&Keluar)

	c.JSON(http.StatusOK, gin.H{"Hapus Data": true})
}
