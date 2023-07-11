package model

type Hello struct {
	ID   uint `gorm:"primary_key"`
	Name string
	Age  uint
}
