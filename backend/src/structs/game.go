package structs

type Game struct {
	Id        int `json:"id"`
	Player1Id int `json:"player1Id"`
	Player2Id int `json:"player2Id"`
	Winner    int `json:"winner"`
	State     string `json:"state"` //active, inactive
	CreatedAt string `json:"createdAt"` 
	UpdatedAt string `json:"updatedAt"`
}

type HistoryGame struct {
	Id       int    `json:"id"`
	GameId   int    `json:"gameUuid"`
	PlayerId int    `json:"playerId"`
	Move     string `json:"move"`
	Duration int    `json:"duration"`
}
