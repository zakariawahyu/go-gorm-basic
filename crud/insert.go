package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"gorm.io/gorm"
)

func InsertProduct(db *gorm.DB, product *entity.Products) {
	result := db.Create(product)

	if result.Error != nil {
		panic("Gagal melakukan insert data produk")
		return
	}
}

func InsertOrder(db *gorm.DB, order *entity.Orders) {
	result := db.Create(order)

	if result.Error != nil {
		panic("Gagal melakukan insert data order")
		return
	}
}

func InsertOrderDetail(db *gorm.DB, orderDetail *[]entity.OrderDetails) {
	result := db.Create(orderDetail)

	if result.Error != nil {
		panic("Gagal melakukan insert data detail order")
		return
	}
}
