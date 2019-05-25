package db

import (
	"fmt"
	"go-pkg/conf"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Setup() {
	DB = open(conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Name, conf.Database.Charset, conf.Database.ParseTime, conf.Database.Loc, conf.Database.Type, conf.Database.SingularTable)
}

func Close() {
	defer DB.Close()
}

func GetDB(section string) *gorm.DB {
	s := conf.Section(section)
	singularTable, _ := strconv.ParseBool(s["SingularTable"])
	return open(s["User"], s["Password"], s["Host"], s["Port"], s["Name"], s["Charset"], s["ParseTime"], s["Loc"], s["Type"], singularTable)
}

func open(user, password, host, port, name, charset, parseTime, loc, dialect string, singular bool) *gorm.DB {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", user, password, host, port, name, charset, parseTime, loc)
	db, err := gorm.Open(dialect, connect)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		panic(err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
		panic(db.Error)
	}

	db.SingularTable(singular)
	return db
}
