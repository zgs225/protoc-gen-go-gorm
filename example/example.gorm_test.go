package example

import (
	"testing"

	"github.com/k0kubun/pp/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newConn() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/hello_db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return db
}

func TestQuery(t *testing.T) {
	m := HelloModel{}
	db := newConn()
	db.Joins("Things").First(&m, 5)

	pp.Println(m)
}

func TestInsert(t *testing.T) {
	m := HelloModel{
		User:   "Zhang San",
		Status: Status_Normal,
		Things: []*ThingModel{
			{
				Name: "牛奶",
			},
			{
				Name: "咖啡",
			},
		},
	}
	db := newConn()
	db.Create(&m)
}
