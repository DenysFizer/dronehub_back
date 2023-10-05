package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Admin    bool   `json:"admin" gorm:"default:false"`
}
type Drone struct {
	Id         uint   `json:"id"`
	Creator    string `json:"creator"`
	Name       string `json:"name" gorm:"unique"`
	Text       string `json:"text"`
	Maxspeed   int    `json:"maxspeed"`
	Batterycap int    `json:"batterycap"`
	Height     uint   `json:"height"`
	Range      uint   `json:"range"`
	Fpv        bool   `json:"fpv"`
	Vizor      bool   `json:"vizor"`
	Imgpath    string `json:"imgpath"`
}
type Comments struct {
	Id       uint   `json:"id"`
	DroneId  uint   `json:"droneid"`
	UserName string `json:"username"`
	Comtext  string `json:"comtext"`
}
