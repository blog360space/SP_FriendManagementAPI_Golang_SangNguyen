package persistence

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"spapp/src/models/domain"
)

var DbContext = UseMySql()

func UseMySql() *gorp.DbMap {
	driver := "mysql"
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, dbname)
	db, err := sql.Open(driver, connectionString)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(domain.UserDomain{}, "User").SetKeys(true, "ID")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}