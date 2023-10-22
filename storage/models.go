package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Password string
	Admin    bool
}

type Application struct {
	gorm.Model
	Name   string
	Token  string
	UserID uint `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
}

type Client struct {
	gorm.Model
	Name    string
	Token   string `gorm:"unique;not null"`
	Enabled bool   `gorm:"not null;default:true"`
	UserID  uint   `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	Chats   []Chat
}

type Chat struct {
	gorm.Model
	Name     string
	Enabled  bool `gorm:"not null;default:true"`
	ChatID   uint `gorm:"unique;not null"`
	ClientID uint `gorm:"foreignkey:ClientID;constraint:OnDelete:CASCADE"`
}

type ClientResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Enabled string `json:"enabled"`
}
