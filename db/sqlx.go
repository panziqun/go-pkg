package db

import (
	"fmt"

	"github.com/laughmaker/go-pkg/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NoData no data in result
const NoData = "sql: no rows in result set"

// DBX export sqlx.DB
var DBX *sqlx.DB

// SetupDBX config sqlx db
func SetupDBX() {
	DBX = openDBX(conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Name, conf.Database.Loc)
}

// CloseDBX close sql db
func CloseDBX() {
	defer DBX.Close()
}

// GetDBX access sql db
func GetDBX(section string) *sqlx.DB {
	s := conf.Section(section)
	return openDBX(s["User"], s["Password"], s["Host"], s["Name"], s["Loc"])
}

func openDBX(user, password, host, name, loc string) *sqlx.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", user, password, host, name, loc)
	db, err := sqlx.Connect("mysql", dns)

	if err != nil {
		panic(err)
	}

	return db
}
