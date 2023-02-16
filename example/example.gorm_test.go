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
	db.First(&m, 2)

	pp.Println(m)
}

func TestInsert(t *testing.T) {
	m := HelloModel{
		User: "Zhang San",
		LogoFile: &LogoFile{
			Url: "hello world",
		},
	}
	db := newConn()
	db.Create(&m)
}
