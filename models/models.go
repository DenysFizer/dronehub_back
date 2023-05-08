package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Admin    bool   `json:"admin" gorm:"default:false"`
}
type Drone struct {
	Id     uint
	Name   string
	Weight uint
	Range  uint
}
