package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/laughmaker/go-pkg/conf"
)

var DB *gorm.DB

func Setup() {
	var err error
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.Charset,
		conf.Database.ParseTime,
		conf.Database.Loc)
	DB, err = gorm.Open(conf.Database.Type, connect)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	DB.SingularTable(conf.Database.SingularTable)

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}

func Close() {
	defer DB.Close()
}
