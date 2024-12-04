package Model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint   `gorm:"not null"` // 发送者ID
	ReceiverID uint   `gorm:"not null"` // 接收者ID
	Content    string `gorm:"not null"`
}
