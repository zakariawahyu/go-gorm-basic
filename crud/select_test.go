package crud

import (
	"fmt"
	"github.com/zakariawahyu/go-gorm-basic/config"
	"testing"
)

func TestShowAllProducts(t *testing.T) {
	result := ShowAllProducts(config.ConnectDB())

	for _, val := range result {
		fmt.Println(val)
	}
}

func TestShowProducts(t *testing.T) {
	result := ShowProducts(config.ConnectDB(), "002")

	if result.KodeProduk == "" {
		t.Error("Produk tidak ditemukan")
	}
	fmt.Println(result)
}

func TestShowAllOrders(t *testing.T) {
	result := ShowAllOrders(config.ConnectDB())

	for _, val := range result {
		fmt.Println(val)
	}
}

func TestShowOrders(t *testing.T) {
	result := ShowOrders(config.ConnectDB(), 1)

	for _, val := range result {
		fmt.Println(val)
	}
}
