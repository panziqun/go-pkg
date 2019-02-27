package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/laughmaker/go-pkg/config"
)

var DB *gorm.DB

func Setup() {
	var err error
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.Name,
		config.DatabaseConfig.Charset,
		config.DatabaseConfig.ParseTime,
		config.DatabaseConfig.Loc)
	DB, err = gorm.Open(config.DatabaseConfig.Type, connect)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	DB.SingularTable(config.DatabaseConfig.SingularTable)

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}

func Close() {
	defer DB.Close()
}
