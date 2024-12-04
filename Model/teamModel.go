package Model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name           string `gorm:"unique;not null"`
	ApprovalStatus string `gorm:"not null"`
	Details        Activity
	Leader         string `gorm:"not null"`
	Member         []User
}

type Activity struct {
	Description     string `gorm:"default:'Welcome to join the activity!'"`
	StartTime       string `gorm:"not null"`
	Name            string `gorm:"not null"`
	Goal            string `gorm:"not null"`
	Destination     string `gorm:"not null"`
	SpecificRequest string
	HumanRequest    uint   `gorm:"default:5"`
	Sort            string `gorm:"not null"`
	// 0: 未开始 1: 进行中 2: 已结束
	Status uint `gorm:"default:0"`
	// 0: 短期 1: 长期保留
	Reserved uint `gorm:"default:0"`
}
