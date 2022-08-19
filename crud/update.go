package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"gorm.io/gorm"
)

func UpdateProduct(db *gorm.DB, kodeProduk string, dataProduct entity.Products) {
	results := db.Where("kode_produk = ?", kodeProduk).Updates(&dataProduct)

	if results.Error != nil {
		panic("Gagal update data produk")
	}
}

func UpdateOrder(db *gorm.DB, idOrder int, dataOrder entity.Orders) {
	results := db.Where("id_order = ?", idOrder).Updates(&dataOrder)

	if results.Error != nil {
		panic("Gagal update data produk")
	}
}
