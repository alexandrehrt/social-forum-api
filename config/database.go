package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	driver := viper.Get("DB_DRIVER")
	host := viper.Get("DB_HOST")
	port := viper.Get("DB_PORT")
	user := viper.Get("DB_USER")
	pass := viper.Get("DB_PASSWORD")
	name := viper.Get("DB_NAME")

	var err error

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driver, user, pass, host, port, name)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}
