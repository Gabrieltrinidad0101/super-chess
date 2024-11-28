package model

import (
	"backend/src/database"
	"backend/src/structs"
)

type ModelGame struct{}

var db = database.GetConnection()

func (m ModelGame) Insert(game *structs.Game) {
	db.Create(game)
}
func (m ModelGame) FindByUuid(uuid string) *structs.Game {
	var game structs.Game
	db.Where("uuid = ?", uuid).First(&game)
	return &game
}
func (m ModelGame) FindByPlayerId(playerId int) *structs.Game {
	var game structs.Game
	db.Where("player_id = ?", playerId).First(&game)
	return &game
}
func (m ModelGame) Update(game *structs.Game) {
	db.Updates(game)
}

func (m ModelGame) FindAll() *[]structs.Game {
	var games []structs.Game
	db.Find(&games).Where("status != finished")
	return &games
}

func (m ModelGame) Delete(gameUuid string) {
	db = db.Debug()
	db.Where("uuid = ?", gameUuid).Delete(&structs.Game{})
}
