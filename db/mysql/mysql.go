package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/RyanDervisevic/apiMarvel/db"
	"github.com/RyanDervisevic/apiMarvel/db/sqlite"
	"github.com/RyanDervisevic/apiMarvel/model"
)

type MySQL = sqlite.SQLite

func New(dbName, heroes, pass, port string) *db.Storage {
	dsn := fmt.Sprintf("%v:%v@tcp(localhost:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", heroes, pass, port, dbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.Heroes{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		Heroes: &MySQL{
			Conn: conn,
		},
	}
}
