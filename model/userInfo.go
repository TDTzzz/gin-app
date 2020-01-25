package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfo struct {
	Uid        int       `gorm:"primary_key;not_null;column:uid"`
	UserName   string    `gorm:"default:null;size:64;column:username"`
	Department string    `gorm:"default:null;size:64;column:department"`
	Created    time.Time `gorm:"default:null;column:created"`
}

type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
}

//func (User) TableName() string {
//	return "test_user"
//}


func (UserInfo) TableName() string {
	return "userinfo"
}
