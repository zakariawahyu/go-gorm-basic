package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/config"
	"github.com/zakariawahyu/go-gorm-basic/entity"
	"testing"
	"time"
)

func TestInsertProduct(t *testing.T) {
	connection := config.ConnectDB()

	product := entity.Products{
		KodeProduk: "001",
		NamaProduk: "Indomie Goreng",
		Stok:       100,
		Harga:      2500,
	}
	InsertProduct(connection, &product)

	product = entity.Products{
		KodeProduk: "002",
		NamaProduk: "Sabun LUX",
		Stok:       200,
		Harga:      1000,
	}
	InsertProduct(connection, &product)
}

func TestInsertOrder(t *testing.T) {
	connection := config.ConnectDB()

	order := entity.Orders{
		TanggalOrder:  time.Now(),
		PaymentMethod: "Cash",
		Total:         30000,
	}

	InsertOrder(connection, &order)

	details := []entity.OrderDetails{
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "001",
			},
			Qty: 10,
		},
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "002",
			},
			Qty: 5,
		},
	}

	InsertOrderDetail(connection, &details)
}
