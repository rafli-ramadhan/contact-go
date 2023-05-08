package client

import (
	"fmt"
	"log"
	"time"

	"contact-go/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conDB 			  = config.GetConf()
	connString string = ""
)

type dbOption struct {
	Database string
}

func GetDBConnection(dbSelected string) dbOption {
	return dbOption{
		Database: dbSelected,
	}
}

func (dbOpt dbOption) GetMysqlConnection() (db *sql.DB, err error) {
	var driver string
	if dbOpt.Database == "mysql" {
		driver = "mysql"
		// "username:password@tcp(host:port)/database_name"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v", 
			conDB.Mysqlconf.Username, 
			conDB.Mysqlconf.Password, 
			conDB.Mysqlconf.Host, 
			conDB.Mysqlconf.Port, 
			conDB.Mysqlconf.Database,
		)
	}

	db, err = sql.Open(driver, connString)
	if err != nil {
		log.Print(err)
		return
	}

	log.Printf("Running mysql on %s on port %s\n", conDB.Mysqlconf.Host, conDB.Mysqlconf.Port)
	
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(10*time.Minute)
	db.SetConnMaxLifetime(60*time.Minute)

	return
}

func (dbOpt dbOption) GetMysqlGormConnection() (*gorm.DB, error) {
	var connString string
	if dbOpt.Database == "mysql-gorm" {
		// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
			conDB.Mysqlconf.Username, 
			conDB.Mysqlconf.Password, 
			conDB.Mysqlconf.Host, 
			conDB.Mysqlconf.Port, 
			conDB.Mysqlconf.Database,
		)
	}

	// customize using mysql.New()
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DriverName: "mysql",
				DSN:		connString,
			},
		), &gorm.Config{},
	)

	// db, err := gorm.Open(
	// 	mysql.Open(connString), &gorm.Config{
	// 		SkipDefaultTransaction: true,
	// 		PrepareStmt:			true,
	// 	},
	// )
	if err != nil {
		log.Print(err)
		return nil, err
	}

	db.Begin()

	log.Printf("Running mysql on %s on port %s\n", conDB.Mysqlconf.Host, conDB.Mysqlconf.Port)

	return db, nil
}