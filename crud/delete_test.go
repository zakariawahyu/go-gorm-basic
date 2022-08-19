package crud

import (
	"github.com/zakariawahyu/go-gorm-basic/config"
	"testing"
)

func TestDeleteOrderDetails(t *testing.T) {
	DeleteOrderDetails(config.ConnectDB(), 3)
}
