package model

import "time"

type HouseInfo struct {
	Id          uint    `gorm:"primary_key;not_null;column:id;AUTO_INCREMENT"`
	Title       string  `gorm:";not_null;type:varchar(50)"`
	UnitPrice   uint    `gorm:"default:0;not_null;size:12"`
	TotalPrice  float64 `gorm:"default:null;type:varchar(50);not_null;default:0"`
	HouseId     string  `gorm:"type:varchar(50);not_null;default:''"`
	Community   string  `gorm:"not_null;type:varchar(50);default:''"`
	CommunityId string  `gorm:"not_null;type:varchar(50);default:''"`
	HouseDetail
	CreateAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type HouseDetail struct {
	HouseType string  `gorm:"not_null;type:varchar(50);default:''"`
	Area      float64 `gorm:"not_null;default:0.00"`
	Toward    string  `gorm:"not_null;type:varchar(50);default:''"`
	Level     string  `gorm:"not_null;type:varchar(50);default:''"`
	Floor     string  `gorm:"not_null;type:varchar(50);default:''"`
	BuildYear uint    `gorm:"not_null;default:0;size:11"`
	BuildType string  `gorm:"not_null;type:varchar(50);default:''"`
}

func (HouseInfo) TableName() string {
	return "houseInfo"
}
