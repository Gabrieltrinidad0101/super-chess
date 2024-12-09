package services

import (
	"backend/src/structs"
	"backend/src/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var eventBus = utils.GetEventBus()

type PlayerModel interface {
	Insert(*structs.Player)
	Find(*structs.Player) structs.Player
	FindById(uint) *structs.Player
	FindByGame(string) []structs.Player
	FindAll() []structs.Player
	SetGame(playerId uint, GameUuid string)
	Winner(playerId uint)
}

type Token interface {
	CreateToken(id uint) (string, error)
}

func NewPlayer(playerModel PlayerModel, token Token) *ServicePlayer {
	return &ServicePlayer{
		PlayerModel: playerModel,
		Token:       token,
	}
}

type ServicePlayer struct {
	PlayerModel
	Token
}

type Response struct {
	StatusCode int
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

var validate = validator.New()

func (u *ServicePlayer) Login(player *structs.Player) Response {
	if err := validate.Struct(player); err != nil {
		return Response{
			StatusCode: 402,
			Message:    "Name or Password  are required",
		}
	}
	playerExist := u.Find(player)
	if playerExist.Name == "" {
		return Response{
			StatusCode: 402,
			Message:    "Name or Password  are incorrect",
		}
	}
	token, error := u.CreateToken(playerExist.Id)

	if error != nil {
		return Response{
			Message:    "Error generating the access token please try later",
			StatusCode: 500,
		}
	}

	return Response{
		Data:       token,
		StatusCode: 200,
	}
}

func (u *ServicePlayer) Register(player *structs.Player) Response {
	if err := validate.Struct(player); err != nil {
		return Response{
			StatusCode: 402,
			Message:    "Name and Password are required",
		}
	}
	playerExist := u.Find(player)
	if playerExist.Name != "" {
		return Response{
			StatusCode: 402,
			Message:    "The user exist",
		}
	}

	u.Insert(player)

	token, error := u.CreateToken(player.Id)
	if error != nil {
		return Response{
			Message:    "Error generating the access token please try later",
			StatusCode: 500,
		}
	}

	return Response{
		Data:       token,
		StatusCode: 200,
	}

}

func (u *ServicePlayer) PlayersByGame(gameId string) Response {
	players := u.FindByGame(gameId)
	return Response{
		StatusCode: 200,
		Message:    players,
	}
}

func (u *ServicePlayer) FindAllPlayers() Response {
	players := u.PlayerModel.FindAll()
	return Response{
		StatusCode: 200,
		Message:    players,
	}
}

func (u *ServicePlayer) SetGame(stringPlayerId, stringGameUuid string) Response {
	playerId, err := strconv.ParseUint(stringPlayerId, 10, 64)
	if err != nil {
		response := Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid player id",
		}
		return response
	}

	// if stringGameUuid != "" {
	// 	player := u.PlayerModel.FindById(uint(playerId))
	// 	if player.GameUuid != "" {
	// 		return Response{
	// 			StatusCode: http.StatusBadRequest,
	// 			Message:    "You are already into game",
	// 		}
	// 	}
	// }
	u.PlayerModel.SetGame(uint(playerId), stringGameUuid)
	return Response{
		StatusCode: 200,
		Message:    "OK",
	}
}

func (u *ServicePlayer) Winner(stringPlayerId string) Response {
	playerId, err := strconv.ParseUint(stringPlayerId, 10, 64)
	if err != nil {
		response := Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid player id",
		}
		return response
	}
	u.PlayerModel.Winner(uint(playerId))
	return Response{
		StatusCode: 200,
		Message:    "OK",
	}
}

func Init() {

}
