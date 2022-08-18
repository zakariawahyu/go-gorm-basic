package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	username := "postgres"
	password := "masukdb"
	database := "golang_gorm"
	port := "5432"

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		panic("Can't connect database")
	}
}
