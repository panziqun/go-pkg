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
		conf.DatabaseConf.User,
		conf.DatabaseConf.Password,
		conf.DatabaseConf.Host,
		conf.DatabaseConf.Port,
		conf.DatabaseConf.Name,
		conf.DatabaseConf.Charset,
		conf.DatabaseConf.ParseTime,
		conf.DatabaseConf.Loc)
	DB, err = gorm.Open(conf.DatabaseConf.Type, connect)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	DB.SingularTable(conf.DatabaseConf.SingularTable)

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}

func Close() {
	defer DB.Close()
}
