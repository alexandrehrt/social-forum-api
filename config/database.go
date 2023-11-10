package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := "postgres://ccigfiia:U4igJ_ya-kpWXcCoBOjCqzqDjq_rb2uC@stampy.db.elephantsql.com/ccigfiia"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}
