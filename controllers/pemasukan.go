package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hadi.com/models"
)

type PemasukanInput struct {
	Id_income string `json:"id_income"`
	Tanggal   string `json:"tanggal"`
	Jumlah    int    `json:"jumlah"`
	Sumber    string `json:"sumber"`
}

//Tampil data pemasukan
func PemasukanTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var masuk []models.Income
	db.Find(&masuk)
	c.JSON(http.StatusOK, gin.H{"data": masuk})
}

// Tambah data pemasukan
func PemasukanTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput PemasukanInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	masuk := models.Income{
		Id_income: dataInput.Id_income,
		Tanggal:   dataInput.Tanggal,
		Jumlah:    dataInput.Jumlah,
		Sumber:    dataInput.Sumber,
	}

	db.Create(&masuk)

	c.JSON(http.StatusOK, gin.H{"data": masuk})
}

// Ubah data Pemasukan
func PemasukanUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var masuk models.Income
	if err := db.Where("id_income = ?", c.Param("id_income")).First(&masuk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Pemasukan tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput PemasukanInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&masuk).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": masuk})
}

// Hapus data Pemasukan
func PemasukanHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var masuk models.Income
	if err := db.Where("id_income = ?", c.Param("id_income")).First(&masuk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Pemasukan tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&masuk)

	c.JSON(http.StatusOK, gin.H{"Hapus data": true})
}
