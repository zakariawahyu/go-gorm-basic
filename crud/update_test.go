package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/config"
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"testing"
)

func TestUpdateProduct(t *testing.T) {
	product := entity.Products{NamaProduk: "Indomie Bawang"}

	UpdateProduct(config.ConnectDB(), "001", product)
}

func TestUpdateOrder(t *testing.T) {
	order := entity.Orders{PaymentMethod: "COD"}

	UpdateOrder(config.ConnectDB(), 1, order)
}
