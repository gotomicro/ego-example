// +build egoctl
package main

type User struct {
	Uid      int    `gorm:"AUTO_INCREMENT" json:"id" dto:"" ego:"primary_key"` // id
	UserName string `gorm:"not null" json:"userName" dto:""`                   // 昵称
}
