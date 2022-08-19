package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"gorm.io/gorm"
)

func ShowAllProducts(db *gorm.DB) []entity.Products {
	var product []entity.Products

	result := db.Find(&product)

	if result.Error != nil {
		panic("Gagal menampilkan data produk")
	}

	return product
}

func ShowProducts(db *gorm.DB, kodeProduk string) entity.Products {
	var product entity.Products

	result := db.Where("kode_produk = ?", kodeProduk).First(&product)

	if result.Error != nil {
		panic("Gagal menampilkan data produk")
	}

	return product
}

func ShowAllOrders(db *gorm.DB) []entity.OrderDetails {
	var orders []entity.OrderDetails

	//result := db.Joins("Products").Joins("Orders").Find(&orders)
	//atau menggunakan preload seperti ini
	result := db.Joins("join products as pd ON pd.kode_produk = order_details.kode_produk").Preload("Products").Joins("Orders").Find(&orders)

	if result.Error != nil {
		panic("Gagal menampilkan data order")
	}

	return orders
}

func ShowOrders(db *gorm.DB, idOrder int) []entity.OrderDetails {
	var orders []entity.OrderDetails

	result := db.Joins("Products").Joins("Orders").Where(&entity.OrderDetails{OrdersIdOrder: idOrder}).Find(&orders)

	if result.Error != nil {
		panic("Gagal menampilkan data order")
	}

	return orders
}
