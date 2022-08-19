package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"gorm.io/gorm"
)

func DeleteOrderDetails(db *gorm.DB, idOrderDetail int) {
	result := db.Where("id_order_detail = ?", idOrderDetail).Delete(&entity.OrderDetails{})

	if result.Error != nil {
		panic("Gagal delete order detail")
	}
}
