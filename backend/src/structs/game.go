package structs

type Game struct {
	Uuid      string `json:"uuid"`
	Player1Id int    `json:"player1Id"`
	Player2Id int    `json:"player2Id"`
}
