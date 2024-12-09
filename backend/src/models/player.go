package model

import (
	"backend/src/structs"
	"fmt"

	"gorm.io/gorm"
)

type PlayerModel interface {
	Find(playerToSearch *structs.Player) structs.Player
	FindById(playerId uint) *structs.Player
	Insert(player *structs.Player)
	FindByGame(GameUuid string) []structs.Player
	SetGame(playerId uint, gameUuid string)
	FindAll() []structs.Player
	Winner(playerId uint)
}

type playerModel struct{}

func (u playerModel) Find(playerToSearch *structs.Player) structs.Player {
	var player structs.Player
	db.Where("name = ?", playerToSearch.Name).First(&player)
	return player
}

func (u playerModel) FindById(playerId uint) *structs.Player {
	var player structs.Player
	db.Where("id = ?", playerId).First(&player)
	return &player
}

func (u playerModel) Insert(player *structs.Player) {
	db.Create(player)
}

func (u playerModel) FindByGame(GameUuid string) []structs.Player {
	var player []structs.Player
	db.Model(&player).Where("game_uuid = ?", GameUuid)
	return player
}

func (u playerModel) SetGame(playerId uint, gameUuid string) {
	var player []structs.Player
	db.Model(&player).Where("id = ?", playerId).Update("game_uuid", gameUuid)
}

func (u playerModel) FindAll() []structs.Player {
	var player []structs.Player
	db.Find(&player).Order("victories DESC")
	return player
}

func (u playerModel) Winner(playerId uint) {
	err := db.Model(structs.Player{}).Where("id = ?", playerId).UpdateColumn("victories", gorm.Expr("victories + ?", 1))
	fmt.Print(err)
}

func NewPlayerModel() playerModel {
	return playerModel{}
}
