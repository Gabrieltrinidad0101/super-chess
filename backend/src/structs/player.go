package structs

type Player struct {
	Name      string `json:"name" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Id        uint   `json:"id"`
	State     string `json:"state"` //active, inactive, waiting
	Victories int    `json:"victories"gorm:"default:0"`
}
