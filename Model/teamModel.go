package Model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name    string     `gorm:"unique;not null"`
	Details []Activity `gorm:"many2many:team_activity;"`
	Leader  string     `gorm:"not null"`
	Member  []User     `gorm:"many2many:user_team;"`
}

type Activity struct {
	gorm.Model
	TeamID          []Team `gorm:"many2many:team_activity;"`
	Description     string `gorm:"default:'Welcome to join the activity!'"`
	StartTime       string `gorm:"not null"`
	ACTName         string `gorm:"not null"`
	Goal            string `gorm:"not null"`
	Destination     string `gorm:"not null"`
	SpecificRequest string
	HumanRequest    uint   `gorm:"default:5"`
	Sort            string `gorm:"not null"`
	ActivityLeader  string `gorm:"not null"`
	// 0: 未开始 1:已结束
	Status uint `gorm:"default:0"`
	// 0: 短期 1: 长期保留
	Reserved uint `gorm:"default:1"`
}

type SelfInstruction struct {
	UserInfo     User
	Introduction string `json:"introduction,omitempty"`
}
