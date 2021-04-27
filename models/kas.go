package models

type Income struct {
	Id_income string `json:"id_income" gorm:"primary_key"`
	Tanggal   string `json:"tanggal"`
	Jumlah    int    `json:"jumlah"`
	Sumber    string `json:"sumber"`
}

type Spending struct {
	Id_spending string `json:"id_spending" gorm:"primary_key"`
	Tanggal     string `json:"tanggal"`
	Jumlah      int    `json:"jumlah"`
	Sumber      string `json:"sumber"`
}
