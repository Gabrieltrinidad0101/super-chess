package structs

type Player struct {
	Name      string `json:"name" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Id        uint   `json:"id"`
	GameUuid  string `json:"gameuuid"`
	Victories int    `json:"victories"gorm:"default:0"`
}
