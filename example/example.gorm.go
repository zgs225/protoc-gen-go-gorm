// Code generated by protoc-gen-go-gorm. DO NOT EDIT.

package example

import (
	sql "database/sql"
	pq "github.com/lib/pq"
	time "time"
)

// 测试用结构

type HelloModel struct {
	Id        uint64
	CreatedAt time.Time
	User      string `gorm:"column:user_name"`
	LogoFile  string `gorm:"column:logo_file2;serializer:json"`
	UpdatedAt sql.NullTime
	Tags      pq.StringArray `gorm:"type:text[]"`
	Status    Status
}

func (m HelloModel) TableName() string {
	return "hello_table"
}
